package dynamodb

import (
	"context"
	"net/url"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	smithyendpoints "github.com/aws/smithy-go/endpoints"

	"github.com/openkrafter/anytore-backend/logger"
)

type DynamoDBService interface {
	Query(ctx context.Context, input *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
	PutItem(ctx context.Context, input *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	GetItem(ctx context.Context, input *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	DeleteItem(ctx context.Context, input *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error)
	UpdateItem(ctx context.Context, input *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)
	Scan(ctx context.Context, input *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error)
}

type resolverV2 struct{}

func (*resolverV2) ResolveEndpoint(ctx context.Context, params dynamodb.EndpointParameters) (
	smithyendpoints.Endpoint, error,
) {
	// u, err := url.Parse("http://localhost:8000")
	u, err := url.Parse(os.Getenv("LOCAL_DYNAMODB_ENDPOINT"))
	if err != nil {
		return smithyendpoints.Endpoint{}, err
	}
	return smithyendpoints.Endpoint{
		URI: *u,
	}, nil
}

type DynamoDBClient struct {
	db *dynamodb.Client
}

func NewDynamoDBClient() (*DynamoDBClient, error) {
	logger.Logger.Debug("Init DynamoDB client.")

	if os.Getenv("DYNAMODB") == "local" {
		// Set dummy credentials for local DynamoDB
		os.Setenv("AWS_ACCESS_KEY_ID", "dummy")
		os.Setenv("AWS_SECRET_ACCESS_KEY", "dummy")
		os.Setenv("AWS_SESSION_TOKEN", "dummy")
	}

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		logger.Logger.Error("Load aws config error.", logger.ErrAttr(err))
		return nil, err
	}

	var dClient *DynamoDBClient
	if os.Getenv("DYNAMODB") == "aws" {
		logger.Logger.Debug("Running in release mode.")
		dClient = &DynamoDBClient{dynamodb.NewFromConfig(cfg)}
	} else if os.Getenv("DYNAMODB") == "local" {
		dClient = &DynamoDBClient{dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
			o.EndpointResolverV2 = &resolverV2{}
		})}
	}

	return dClient, nil
}

func (dClient *DynamoDBClient) Query(ctx context.Context, input *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	return dClient.db.Query(ctx, input, optFns...)
}

func (dClient *DynamoDBClient) PutItem(ctx context.Context, input *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	return dClient.db.PutItem(ctx, input, optFns...)
}

func (dClient *DynamoDBClient) GetItem(ctx context.Context, input *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	return dClient.db.GetItem(ctx, input, optFns...)
}

func (dClient *DynamoDBClient) DeleteItem(ctx context.Context, input *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error) {
	return dClient.db.DeleteItem(ctx, input, optFns...)
}

func (dClient *DynamoDBClient) UpdateItem(ctx context.Context, input *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error) {
	return dClient.db.UpdateItem(ctx, input, optFns...)
}

func (dClient *DynamoDBClient) Scan(ctx context.Context, input *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
	return dClient.db.Scan(ctx, input, optFns...)
}
