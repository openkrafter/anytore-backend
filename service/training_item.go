package service

import (
	"context"
	"errors"
	"log/slog"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go"
	"github.com/openkrafter/anytore-backend/customerror"
	"github.com/openkrafter/anytore-backend/logger"
	"github.com/openkrafter/anytore-backend/model"
)

func GetTraningItems(userId int) ([]*model.TrainingItem, error) {
	basics, err := NewTableBasics("TrainingItem")
	if err != nil {
		logger.Logger.Error("DynamoDB client init error.", logger.ErrAttr(err))
		return nil, err
	}

	keyCond := expression.Key("UserId").Equal(expression.Value(userId))
	proj := expression.NamesList(
		expression.Name("Id"),
		expression.Name("UserId"),
		expression.Name("Name"),
		expression.Name("Type"),
		expression.Name("Unit"),
		expression.Name("Kcal"),
	)

	builder := expression.NewBuilder().
		WithKeyCondition(keyCond).
		WithProjection(proj)
	expr, err := builder.Build()
	if err != nil {
		logger.Logger.Error("Failed to get TrainingItems.", logger.ErrAttr(err))
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		TableName:                 &basics.TableName,
		IndexName:                 aws.String("UserIdIndex"),
	}

	response, err := basics.DynamoDbClient.Query(context.TODO(), queryInput)
	if err != nil {
		logger.Logger.Error("Failed to get TrainingItems.", logger.ErrAttr(err))
		return nil, err
	}

	if len(response.Items) == 0 {
		logger.Logger.Info("No items matched.")
		return nil, nil
	} else {
		logger.Logger.Debug("Query succeeded.")
	}

	var trainingItems []*model.TrainingItem
	for _, item := range response.Items {
		var trainingItem model.TrainingItem
		err = attributevalue.UnmarshalMap(item, &trainingItem)
		if err != nil {
			logger.Logger.Error("Failed to unmarshal response.", logger.ErrAttr(err))
		}
		trainingItems = append(trainingItems, &trainingItem)
	}

	return trainingItems, nil
}

func GetTraningItem(id int, userId int) (*model.TrainingItem, error) {
	logger.Logger.Debug("GetTraningItem process", slog.Int("id", id))

	logger.Logger.Debug("Init DynamoDB client.")
	basics, err := NewTableBasics("TrainingItem")
	if err != nil {
		logger.Logger.Error("DynamoDB client init error.", logger.ErrAttr(err))
		return nil, err
	}

	logger.Logger.Debug("Get TraningItem.")
	keyCond := expression.Key("UserId").Equal(expression.Value(userId))
	filter := expression.Name("Id").Equal(expression.Value(id))
	proj := expression.NamesList(
		expression.Name("Id"),
		expression.Name("UserId"),
		expression.Name("Name"),
		expression.Name("Type"),
		expression.Name("Unit"),
		expression.Name("Kcal"),
	)

	builder := expression.NewBuilder().
		WithKeyCondition(keyCond).
		WithFilter(filter).
		WithProjection(proj)
	expr, err := builder.Build()
	if err != nil {
		logger.Logger.Error("Failed to get TrainingItems.", logger.ErrAttr(err))
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		KeyConditionExpression:    expr.KeyCondition(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		TableName:                 &basics.TableName,
		IndexName:                 aws.String("UserIdIndex"),
	}

	response, err := basics.DynamoDbClient.Query(context.TODO(), queryInput)
	if err != nil {
		logger.Logger.Error("Failed to get TrainingItems.", logger.ErrAttr(err))
		return nil, err
	}
	if len(response.Items) == 0 {
		logger.Logger.Info("No items matched.")
		return nil, nil
	} else {
		logger.Logger.Debug("Query succeeded.")
	}

	trainingItem := new(model.TrainingItem)
	err = attributevalue.UnmarshalMap(response.Items[0], &trainingItem)
	if err != nil {
		logger.Logger.Error("Failed to unmarshal response.", logger.ErrAttr(err))
	}

	logger.Logger.Debug("Success to get TrainingItem.", slog.Any("TrainingItem", trainingItem))

	return trainingItem, nil
}

func GetIncrementId() (int, error) {
	basics, err := NewTableBasics("TrainingItemCounter")
	if err != nil {
		logger.Logger.Error("DynamoDB client init error.", logger.ErrAttr(err))
		return -1, err
	}

	type TrainingItemCounter struct {
		CountKey    string
		CountNumber int
	}

	countKey := map[string]types.AttributeValue{
		"CountKey": &types.AttributeValueMemberS{
			Value: "key",
		},
	}

	updateExpression := "SET CountNumber = CountNumber + :incr"
	conditionExpression := "attribute_exists(CountNumber)"
	incrementExp := map[string]types.AttributeValue{
		":incr": &types.AttributeValueMemberN{
			Value: "1",
		},
	}

	result, err := basics.DynamoDbClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName:                 aws.String("TrainingItemCounter"),
		Key:                       countKey,
		UpdateExpression:          aws.String(updateExpression),
		ExpressionAttributeValues: incrementExp,
		ConditionExpression:       aws.String(conditionExpression),
		ReturnValues:              types.ReturnValueUpdatedNew,
	})
	if err != nil {
		var apiErr smithy.APIError
		if ok := errors.As(err, &apiErr); ok {
			if apiErr.ErrorCode() == "ConditionalCheckFailedException" {
				logger.Logger.Info("No item in TrainingItemCounter table, put initial item.", logger.ErrAttr(err))
				_, err = basics.DynamoDbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
					TableName: aws.String("TrainingItemCounter"),
					Item: map[string]types.AttributeValue{
						"CountKey":    &types.AttributeValueMemberS{Value: "key"},
						"CountNumber": &types.AttributeValueMemberN{Value: "0"},
					},
					ConditionExpression: aws.String("attribute_not_exists(CountKey)"),
				})
				if err != nil {
					logger.Logger.Error("Failed to create TrainingItemCounter item.", logger.ErrAttr(err))

				}
			}
			return 1, nil
		}

		logger.Logger.Error("GetIncrementId Failed.", logger.ErrAttr(err))
		return -1, err
	}

	var updatedAttributes TrainingItemCounter
	err = attributevalue.UnmarshalMap(result.Attributes, &updatedAttributes)
	if err != nil {
		logger.Logger.Error("GetIncrementId Failed.", logger.ErrAttr(err))
		return -1, err
	}

	return int(updatedAttributes.CountNumber), nil
}

func CreateTraningItem(input *model.TrainingItem) error {
	basics, err := NewTableBasics("TrainingItem")
	if err != nil {
		logger.Logger.Error("DynamoDB client init error.", logger.ErrAttr(err))
		return err
	}
	av, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}
	_, err = basics.DynamoDbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(basics.TableName),
		Item:      av,
	})
	if err != nil {
		return err
	}

	return nil
}

func UpdateTraningItem(input *model.TrainingItem, userId int) error {
	basics, err := NewTableBasics("TrainingItem")
	if err != nil {
		logger.Logger.Error("DynamoDB client init error.", logger.ErrAttr(err))
		return err
	}

	trainingItem, err := GetTraningItem(input.Id, userId)
	if err != nil {
		return err
	}
	if trainingItem == nil {
		error404 := customerror.NewError404()
		return &error404
	}

	av, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}
	_, err = basics.DynamoDbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(basics.TableName),
		Item:      av,
	})
	if err != nil {
		return err
	}

	return nil
}

func DeleteTraningItem(id int, userId int) error {
	basics, err := NewTableBasics("TrainingItem")
	if err != nil {
		logger.Logger.Error("DynamoDB client init error.", logger.ErrAttr(err))
		return err
	}

	trainingItem, err := GetTraningItem(id, userId)
	if err != nil {
		return err
	}
	if trainingItem == nil {
		error404 := customerror.NewError404()
		return &error404
	}

	deleteInput := &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberN{
				Value: strconv.Itoa(id),
			},
		},
		TableName: aws.String(basics.TableName),
	}

	_, err = basics.DynamoDbClient.DeleteItem(context.TODO(), deleteInput)
	if err != nil {
		return err
	}

	return nil
}
