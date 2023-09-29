package models

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Movie encapsulates data about a movie. Title and Year are the composite primary key
// of the movie in Amazon DynamoDB. Title is the sort key, Year is the partition key,
// and Info is additional data.
type User struct {
	Username             string   `dynamodbav:"user_name,string"`
	GithubUrl            string   `dynamodbav:"github_url,string"`
	Email                string   `dynamodbav:"email,string"`
	Name                 string   `dynamodbav:"name,string"`
	LeetcodeUrl          string   `dynamodbav:"leetcode_url,string"`
	StackOverflowUrl     string   `dynamodbav:"stack_overflow_url,string"`
	ProgrammingLanguages []string `dynamodbav:"programming_languages,stringset"`
	DataStorage          []string `dynamodbav:"data_storage,stringset"`
	Infrastructure       []string `dynamodbav:"infrastructure,stringset"`
}

func (user User) GetKey() map[string]types.AttributeValue {
	username, err := attributevalue.Marshal(user.Username)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"user_name": username}
}
