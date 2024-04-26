package testenvironment

import (
	"context"
	"log"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	anytoreDynamodb "github.com/openkrafter/anytore-backend/dynamodb"
	"github.com/openkrafter/anytore-backend/logger"
	"github.com/openkrafter/anytore-backend/model"
)

type resolverV2 struct{}

func (*resolverV2) ResolveEndpoint(ctx context.Context, params dynamodb.EndpointParameters) (
	smithyendpoints.Endpoint, error,
) {
	u, err := url.Parse("http://localhost:8000")
	if err != nil {
		return smithyendpoints.Endpoint{}, err
	}
	return smithyendpoints.Endpoint{
		URI: *u,
	}, nil
}

func SetupDynamoDbClient() error {
	cfg, err := config.LoadDefaultConfig(context.Background())

	if err != nil {
		logger.Logger.Error("Load aws config error.", logger.ErrAttr(err))
		return err
	}

	anytoreDynamodb.DynamoDbClient = dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.EndpointResolverV2 = &resolverV2{}
	})

	return nil
}

func SetupTraningItemTestData(input *model.TrainingItem) error {
	tableName := "TrainingItem"
	av, err := attributevalue.MarshalMap(input)
	if err != nil {
		return err
	}
	_, err = anytoreDynamodb.DynamoDbClient.PutItem(context.Background(), &dynamodb.PutItemInput{
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
	_, err = anytoreDynamodb.DynamoDbClient.PutItem(context.Background(), &dynamodb.PutItemInput{
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
	deleteItems, err := anytoreDynamodb.DynamoDbClient.Scan(context.Background(), params)
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

		_, err := anytoreDynamodb.DynamoDbClient.DeleteItem(context.Background(), deleteParams)
		if err != nil {
			logger.Logger.Error("Failed to delete item.", logger.ErrAttr(err))
			return
		}
	}

	log.Printf("Deleted all items in table: %s", tableName)
}
