package models

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Movie encapsulates data about a movie. Title and Year are the composite primary key
// of the movie in Amazon DynamoDB. Title is the sort key, Year is the partition key,
// and Info is additional data.
type User struct {
	Username  string `dynamodbav:"user_name,string"`
	GithubUrl string `dynamodbav:"github_url,string"`
}

func (user User) GetKey() map[string]types.AttributeValue {
	username, err := attributevalue.Marshal(user.Username)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"user_name": username}
}

// type UserTable struct {
// 	DynamoDbClient *dynamodb.Client
// 	TableName      string
// }

// func (usertable UserTable) GetUser(username string) (User, error) {
// 	user := User{Username: username}
// 	resp, err := usertable.DynamoDbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
// 		TableName: aws.String(usertable.TableName),
// 		Key:       user.GetKey(),
// 	},
// 	)
// 	out := User{}
// 	if resp.Item != nil {
// 		// var item User
// 		err = attributevalue.UnmarshalMap(resp.Item, &out)
// 		if err != nil {
// 			log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
// 		}
// 		return out, err
// 	} else {
// 		return User{}, err
// 	}
// }

// func GetTestUsersTable() UserTable {
// 	svc := GetDynamoDbClient()
// 	return UserTable{DynamoDbClient: svc, TableName: "test-users"}
// }

// func GetDynamoDbClient() *dynamodb.Client {
// 	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
// 	if err != nil {
// 		log.Fatalf("unable to load SDK config, %v", err)
// 	}

// 	svc := dynamodb.NewFromConfig(cfg)
// 	return svc
// }
