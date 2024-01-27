package service

import (
	"context"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/openkrafter/anytore-backend/logger"
	"github.com/openkrafter/anytore-backend/model"
)

// tmp
func createTrainingItemSample() model.TrainingItem {
	trainingItem := new(model.TrainingItem)
	trainingItem.Id = 1
	trainingItem.UserId = 1
	trainingItem.Name = "running"
	trainingItem.Type = "aerobic"
	trainingItem.Unit = "minute"
	trainingItem.Kcal = 2

	return *trainingItem
}

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
	} else {
		err = attributevalue.UnmarshalMap(response.Item, &trainingItem)
		if err != nil {
			logger.Logger.Error("Failed to unmarshal response.", logger.ErrAttr(err))
		}
		logger.Logger.Debug("Success to get TrainingItem.", slog.Any("TrainingItem", trainingItem))
	}

	return trainingItem, nil
}

func GetTraningItem(id int) model.TrainingItem {
	logger.Logger.Debug("GetTraningItem process", slog.Int("id", id))

	logger.Logger.Debug("Init DynamoDB client.")
	basics, err := NewTableBasics("TrainingItem")
	if err != nil {
		logger.Logger.Error("DynamoDB client init error.", logger.ErrAttr(err))
		return createTrainingItemSample()
	}

	logger.Logger.Debug("Get TraningItem.")
	trainingItem, err := basics.getTrainingItemById(id)
	if err != nil {
		logger.Logger.Error("Get TraningItem error.", logger.ErrAttr(err))
		return createTrainingItemSample()
	}

	return *trainingItem
}
