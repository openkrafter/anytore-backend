package service

import (
	"context"
	"errors"
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

func GetTrainingRecords(ctx context.Context, userId int) ([]*model.TrainingRecord, error) {
	basics, err := NewTableBasics("TrainingRecord")
	if err != nil {
		logger.Logger.Error("DynamoDB client init error.", logger.ErrAttr(err))
		return nil, err
	}

	keyCond := expression.Key("UserId").Equal(expression.Value(userId))
	proj := expression.NamesList(
		expression.Name("Id"),
		expression.Name("UserId"),
		expression.Name("TrainingItemId"),
		expression.Name("Record"),
		expression.Name("Date"),
	)

	builder := expression.NewBuilder().
		WithKeyCondition(keyCond).
		WithProjection(proj)
	expr, err := builder.Build()
	if err != nil {
		logger.Logger.Error("Failed to get TrainingRecords.", logger.ErrAttr(err))
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

	logger.Logger.Debug("Get TrainingRecords.")
	response, err := basics.DynamoDbClient.Query(ctx, queryInput)
	if err != nil {
		logger.Logger.Error("Failed to get TrainingRecords.", logger.ErrAttr(err))
		return nil, err
	}
	if len(response.Items) == 0 {
		logger.Logger.Info("No items matched.")
		return nil, nil
	} else {
		logger.Logger.Debug("Query succeeded.")
	}

	trainingRecords := make([]*model.TrainingRecord, 0)
	for _, i := range response.Items {
		trainingRecord := new(model.TrainingRecord)
		err = attributevalue.UnmarshalMap(i, trainingRecord)
		if err != nil {
			logger.Logger.Error("Failed to unmarshal TrainingRecord.", logger.ErrAttr(err))
			return nil, err
		}
		trainingRecords = append(trainingRecords, trainingRecord)
	}

	return trainingRecords, nil
}

func GetTrainingRecord(ctx context.Context, trainingRecordId int, userId int) (*model.TrainingRecord, error) {
	basics, err := NewTableBasics("TrainingRecord")
	if err != nil {
		logger.Logger.Error("DynamoDB client init error.", logger.ErrAttr(err))
		return nil, err
	}

	keyCond := expression.Key("UserId").Equal(expression.Value(userId))
	filter := expression.Name("Id").Equal(expression.Value(trainingRecordId))
	proj := expression.NamesList(
		expression.Name("Id"),
		expression.Name("UserId"),
		expression.Name("TrainingItemId"),
		expression.Name("Record"),
		expression.Name("Date"),
	)

	builder := expression.NewBuilder().
		WithKeyCondition(keyCond).
		WithFilter(filter).
		WithProjection(proj)
	expr, err := builder.Build()
	if err != nil {
		logger.Logger.Error("Failed to get TrainingRecord.", logger.ErrAttr(err))
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

	logger.Logger.Debug("Get TrainingRecord.")
	response, err := basics.DynamoDbClient.Query(ctx, queryInput)
	if err != nil {
		logger.Logger.Error("Failed to get TrainingRecord.", logger.ErrAttr(err))
		return nil, err
	}
	if len(response.Items) == 0 {
		logger.Logger.Info("No items matched.")
		return nil, nil
	} else {
		logger.Logger.Debug("Query succeeded.")
	}

	trainingRecord := new(model.TrainingRecord)
	err = attributevalue.UnmarshalMap(response.Items[0], trainingRecord)
	if err != nil {
		logger.Logger.Error("Failed to unmarshal TrainingRecord.", logger.ErrAttr(err))
		return nil, err
	}

	return trainingRecord, nil
}

func CreateTrainingRecord(ctx context.Context, trainingRecord *model.TrainingRecord) error {
	basics, err := NewTableBasics("TrainingRecord")
	if err != nil {
		logger.Logger.Error("DynamoDB client init error.", logger.ErrAttr(err))
		return err
	}

	trainingRecord.Id, err = getTrainingRecordIncrementId(ctx)
	if err != nil {
		logger.Logger.Error("Failed to get increment ID.", logger.ErrAttr(err))
		return err
	}

	av, err := attributevalue.MarshalMap(trainingRecord)
	if err != nil {
		logger.Logger.Error("Failed to marshal TrainingRecord.", logger.ErrAttr(err))
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(basics.TableName),
	}

	logger.Logger.Debug("Create TrainingRecord.")
	_, err = basics.DynamoDbClient.PutItem(ctx, input)
	if err != nil {
		logger.Logger.Error("Failed to create TrainingRecord.", logger.ErrAttr(err))
		return err
	}

	return nil
}

func UpdateTrainingRecord(ctx context.Context, trainingRecord *model.TrainingRecord, userId int) error {
	basics, err := NewTableBasics("TrainingRecord")
	if err != nil {
		logger.Logger.Error("DynamoDB client init error.", logger.ErrAttr(err))
		return err
	}

	getTrainingRecordResult, err := GetTrainingRecord(ctx, trainingRecord.Id, userId)
	if err != nil {
		logger.Logger.Error("Failed to get TrainingRecord.", logger.ErrAttr(err))
		return err
	}
	if getTrainingRecordResult == nil {
		error404 := customerror.NewError404()
		return &error404
	}

	av, err := attributevalue.MarshalMap(trainingRecord)
	if err != nil {
		logger.Logger.Error("Failed to marshal TrainingRecord.", logger.ErrAttr(err))
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(basics.TableName),
	}

	logger.Logger.Debug("Update TrainingRecord.")
	_, err = basics.DynamoDbClient.PutItem(ctx, input)
	if err != nil {
		logger.Logger.Error("Failed to update TrainingRecord.", logger.ErrAttr(err))
		return err
	}

	return nil
}

func DeleteTrainingRecord(ctx context.Context, trainingRecordId int, userId int) error {
	basics, err := NewTableBasics("TrainingRecord")
	if err != nil {
		logger.Logger.Error("DynamoDB client init error.", logger.ErrAttr(err))
		return err
	}

	getTrainingRecordResult, err := GetTrainingRecord(ctx, trainingRecordId, userId)
	if err != nil {
		logger.Logger.Error("Failed to get TrainingRecord.", logger.ErrAttr(err))
		return err
	}
	if getTrainingRecordResult == nil {
		error404 := customerror.NewError404()
		return &error404
	}

	deleteInput := &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberN{
				Value: strconv.Itoa(trainingRecordId),
			},
		},
		TableName: aws.String(basics.TableName),
	}

	logger.Logger.Debug("Delete TrainingRecord.")
	_, err = basics.DynamoDbClient.DeleteItem(ctx, deleteInput)
	if err != nil {
		logger.Logger.Error("Failed to delete TrainingRecord.", logger.ErrAttr(err))
		return err
	}

	return nil
}

func getTrainingRecordIncrementId(ctx context.Context) (int, error) {
	basics, err := NewTableBasics("TrainingRecordCounter")
	if err != nil {
		logger.Logger.Error("DynamoDB client init error.", logger.ErrAttr(err))
		return -1, err
	}

	type TrainingRecordCounter struct {
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

	result, err := basics.DynamoDbClient.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 aws.String("TrainingRecordCounter"),
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
				logger.Logger.Info("No item in TrainingRecordCounter table, put initial item.", logger.ErrAttr(err))
				_, err = basics.DynamoDbClient.PutItem(ctx, &dynamodb.PutItemInput{
					TableName: aws.String("TrainingRecordCounter"),
					Item: map[string]types.AttributeValue{
						"CountKey":    &types.AttributeValueMemberS{Value: "key"},
						"CountNumber": &types.AttributeValueMemberN{Value: "0"},
					},
					ConditionExpression: aws.String("attribute_not_exists(CountKey)"),
				})
				if err != nil {
					logger.Logger.Error("Failed to create TrainingRecordCounter item.", logger.ErrAttr(err))

				}
			}
			return 1, nil
		}

		logger.Logger.Error("GetIncrementId Failed.", logger.ErrAttr(err))
		return -1, err
	}

	var updatedAttributes TrainingRecordCounter
	err = attributevalue.UnmarshalMap(result.Attributes, &updatedAttributes)
	if err != nil {
		logger.Logger.Error("GetIncrementId Failed.", logger.ErrAttr(err))
		return -1, err
	}

	return int(updatedAttributes.CountNumber), nil
}
