package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// type User struct {
// 	Username string
// }

func main() {
	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	// Using the Config value, create the DynamoDB client
	svc := GetDynamoDbClient()

	// Build the request with its input parameters
	// tableNames := GetTableNames(svc)
	userTable := UserTable{DynamoDbClient: svc, TableName: "test-users"}
	search_user, err := UserTable.GetUser(userTable, "test_user")
	// fmt.Println("Tables:")
	// for _, tableName := range tableNames {
	// 	fmt.Println(tableName)
	// }
	if err != nil {
		log.Fatalf("failed to get item, %v", err)
	} else {
		fmt.Println("User:")
		fmt.Println(search_user)
	}
}

type UserTable struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

func GetTableNames(svc *dynamodb.Client) []string {
	resp, err := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{
		Limit: aws.Int32(5),
	})
	if err != nil {
		log.Fatalf("failed to list tables, %v", err)
	}
	tableNames := resp.TableNames
	return tableNames
}

func (usertable UserTable) GetUser(username string) (User, error) {
	user := User{username: username}
	resp, err := usertable.DynamoDbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{

		TableName: aws.String("Users"),
		Key:       user.GetKey(),
	},
	)
	if err != nil {
		log.Fatalf("failed to get item, %v", err)
	} else {
		err = attributevalue.UnmarshalMap(resp.Item, &user)
		if err != nil {
			log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		}
	}
	return user, err
}

func GetDynamoDbClient() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := dynamodb.NewFromConfig(cfg)
	return svc
}
