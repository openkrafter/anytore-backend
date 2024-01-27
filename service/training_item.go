package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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
	trainingItem.CreatedAt = time.Now()
	trainingItem.UpdatedAt = time.Now()

	return *trainingItem
}

func (basics *TableBasics) getTrainingItemById(id int) (*model.TrainingItem, error) {
	searchId, err := attributevalue.Marshal(id)
	if err != nil {
		fmt.Println("failed to marshal")
		return nil, err
	}
	searchKey := map[string]types.AttributeValue{"Id": searchId}

	trainingItem := new(model.TrainingItem)
	response, err := basics.DynamoDbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: searchKey, TableName: &basics.TableName,
    })
	if err != nil {
        log.Fatalf("failed to get item, %v", err)
		return nil, err
	} else {
		err = attributevalue.UnmarshalMap(response.Item, &trainingItem)
		if err != nil {
			log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		}
		log.Printf(trainingItem.Name)
		log.Printf(trainingItem.Type)
		log.Printf(trainingItem.Unit)
		log.Printf("TrainingItem: %v", trainingItem)
	}

	return trainingItem, nil
}

func GetTraningItem(id int) model.TrainingItem {
	fmt.Println("trainingItem id: ", id)

	fmt.Println("Connect DynamoDB")
	basics, err := NewTableBasics("TrainingItem")
	if err != nil {
		fmt.Println("connect error: DynamoDB")
		return createTrainingItemSample()
	}

	fmt.Println("Get TraningItem")
	trainingItem, err := basics.getTrainingItemById(id)
	if err != nil {
		fmt.Println("get error: TraningItem")
		return createTrainingItemSample()
	}

	return *trainingItem
}
