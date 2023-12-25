package adapter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBRepository struct {
	DynamoDBClient DynamoDBAPI
	TableName      string
}

type DynamoDBAPI interface {
	DeleteItem(*dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error)
}

func NewDynamoDBRepository(tableName string, client DynamoDBAPI) *DynamoDBRepository {
	return &DynamoDBRepository{
		DynamoDBClient: client,
		TableName:      tableName,
	}
}

func (r *DynamoDBRepository) DeletePointByUser(id string) error {
	key := map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String(id),
		},
	}

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(r.TableName),
		Key:       key,
	}

	_, err := r.DynamoDBClient.DeleteItem(input)
	if err != nil {
		return err
	}

	return nil
}
