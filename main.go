package main

import (
	"businesscardapi/repository"
	"fmt"
	"log"
)

func main() {
	testUserName := "test_user_2"
	userTable := repository.GetTestUsersTable()
	search_user, err := userTable.GetUser(testUserName)

	if err != nil || search_user.Username == "" {
		log.Printf("failed to get user, %v, here's why %v", testUserName, err)
	} else {
		fmt.Println("User:")
		fmt.Println(search_user)
	}
}
