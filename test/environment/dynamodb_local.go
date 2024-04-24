package testenvironment

import (
	"context"
	"net/url"
	"strconv"

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
	cfg, err := config.LoadDefaultConfig(context.TODO())

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

func TeardownTraningItemTestData(id int) error {
	tableName := "TrainingItem"
	deleteInput := &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberN{
				Value: strconv.Itoa(id),
			},
		},
		TableName: aws.String(tableName),
	}

	_, err := anytoreDynamodb.DynamoDbClient.DeleteItem(context.Background(), deleteInput)
	if err != nil {
		return err
	}

	return nil
}
