package testenvironment

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	anytoreDynamodb "github.com/openkrafter/anytore-backend/database/dynamodb"
	"github.com/openkrafter/anytore-backend/logger"
	"github.com/openkrafter/anytore-backend/model"
)

func SetupDynamoDbClient() error {
	os.Setenv("DYNAMODB", "local")
	os.Setenv("LOCAL_DYNAMODB_ENDPOINT", "http://localhost:8000")

	var err error
	anytoreDynamodb.DClient, err = anytoreDynamodb.NewDynamoDBClient()
	if err != nil {
		errMsg := "Failed to init DynamoDBClient."
		logger.Logger.Error(errMsg, logger.ErrAttr(err))
		return err
	}

	return nil
}

func SetupTraningItemTestData(input *model.TrainingItem) error {
	tableName := "TrainingItem"
	av, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}
	_, err = anytoreDynamodb.DClient.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	})
	if err != nil {
		return err
	}

	return nil
}

func SetupTraningItemCounterTestData(input *model.TrainingItem) error {
	tableName := "TrainingItem"
	av, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}
	_, err = anytoreDynamodb.DClient.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	})
	if err != nil {
		return err
	}

	return nil
}

func TeardownTraningItemTestData() error {
	tables := [][]string{
		{"TrainingItem", "Id"},
		{"TrainingItemCounter", "CountKey"},
	}

	for _, table := range tables {
		deleteAllItems(table[0], table[1])
	}

	return nil
}

func deleteAllItems(tableName string, keyName string) {
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	deleteItems, err := anytoreDynamodb.DClient.Scan(context.Background(), params)
	if err != nil {
		logger.Logger.Error("Failed to scan items.", logger.ErrAttr(err))
		return
	}

	for _, item := range deleteItems.Items {
		log.Printf("Delete item: %v", item)
		deleteParams := &dynamodb.DeleteItemInput{
			TableName: aws.String(tableName),
			Key: map[string]types.AttributeValue{
				keyName: item[keyName],
			},
		}

		_, err := anytoreDynamodb.DClient.DeleteItem(context.Background(), deleteParams)
		if err != nil {
			logger.Logger.Error("Failed to delete item.", logger.ErrAttr(err))
			return
		}
	}

	log.Printf("Deleted all items in table: %s", tableName)
}
