package service

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/openkrafter/anytore-backend/logger"
	"github.com/openkrafter/anytore-backend/model"
)

func (basics *TableBasics) getTrainingItemById(id int) (*model.TrainingItem, error) {
	searchId, err := attributevalue.Marshal(id)
	if err != nil {
		logger.Logger.Error("Failed to marshal.", logger.ErrAttr(err))
		return nil, err
	}
	searchKey := map[string]types.AttributeValue{"Id": searchId}

	trainingItem := new(model.TrainingItem)
	response, err := basics.DynamoDbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: searchKey, TableName: &basics.TableName,
    })
	if err != nil {
		logger.Logger.Error("Failed to get TrainingItem.", logger.ErrAttr(err))
		return nil, err
	}
	err = attributevalue.UnmarshalMap(response.Item, &trainingItem)
	if err != nil {
		logger.Logger.Error("Failed to unmarshal response.", logger.ErrAttr(err))
	}
	logger.Logger.Debug("Success to get TrainingItem.", slog.Any("TrainingItem", trainingItem))

	return trainingItem, nil
}

func GetTraningItems() ([]*model.TrainingItem, error) {
	basics, err := NewTableBasics("TrainingItem")
	if err != nil {
		logger.Logger.Error("DynamoDB client init error.", logger.ErrAttr(err))
		return nil, err
	}

	response, err := basics.DynamoDbClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: &basics.TableName,
	})	
	if err != nil {
		logger.Logger.Error("Failed to scan TrainingItem.", logger.ErrAttr(err))
		return nil, err
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

func GetTraningItem(id int) (*model.TrainingItem, error) {
	logger.Logger.Debug("GetTraningItem process", slog.Int("id", id))

	logger.Logger.Debug("Init DynamoDB client.")
	basics, err := NewTableBasics("TrainingItem")
	if err != nil {
		logger.Logger.Error("DynamoDB client init error.", logger.ErrAttr(err))
		return nil, err
	}

	logger.Logger.Debug("Get TraningItem.")
	trainingItem, err := basics.getTrainingItemById(id)
	if err != nil {
		logger.Logger.Error("Get TraningItem error.", logger.ErrAttr(err))
		return nil, err
	}

	return trainingItem, nil
}

func GetIncrementId() (int, error) {
	basics, err := NewTableBasics("TrainingItemCounter")
	if err != nil {
		logger.Logger.Error("DynamoDB client init error.", logger.ErrAttr(err))
		return -1, err
	}

	type TrainingItemCounter struct {
		CountKey string
		CountNumber int
	}

	countKey := map[string]types.AttributeValue{
		"CountKey": &types.AttributeValueMemberS{
			Value: "key",
		},
	}

	updateExpression := "SET CountNumber = CountNumber + :incr"
	incrementExp := map[string]types.AttributeValue{
		":incr": &types.AttributeValueMemberN{
			Value: "1",
		},
	}

	result, err := basics.DynamoDbClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String("TrainingItemCounter"),
		Key: countKey,
		UpdateExpression: aws.String(updateExpression),
		ExpressionAttributeValues: incrementExp,
		ReturnValues: types.ReturnValueUpdatedNew,
	})
	if err != nil {
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

func UpdateTraningItem(input *model.TrainingItem) error {
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

func DeleteTraningItem(id int) error {
	basics, err := NewTableBasics("TrainingItem")
	if err != nil {
		logger.Logger.Error("DynamoDB client init error.", logger.ErrAttr(err))
		return err
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

