package dynamo

import (
	"context"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/joshpauline/dynauth/domain/auth"
)

type dynamoImpl struct {
	client *dynamodb.DynamoDB
}

func NewDynamoImpl(client *dynamodb.DynamoDB) auth.AuthRepository {
	return &dynamoImpl{
		client,
	}
}

func (d *dynamoImpl) GetUser(ctx context.Context) error {
	return nil
}
