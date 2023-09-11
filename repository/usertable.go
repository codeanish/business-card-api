package repository

import (
	"businesscardapi/models"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type UserTable struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func (usertable UserTable) GetUser(username string) (models.User, error) {
	user := models.User{Username: username}
	resp, err := usertable.DynamoDbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(usertable.TableName),
		Key:       user.GetKey(),
	},
	)
	out := models.User{}
	if resp.Item != nil {
		// var item User
		err = attributevalue.UnmarshalMap(resp.Item, &out)
		if err != nil {
			log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		}
		return out, err
	} else {
		return models.User{}, err
	}
}

func GetTestUsersTable() UserTable {
	svc := GetDynamoDbClient()
	return UserTable{DynamoDbClient: svc, TableName: "test-users"}
}
