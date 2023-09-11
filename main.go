package main

import (
	"businesscardapi/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET(":username", getUserDetails)
	router.Run("localhost:8080")
}

func getUserDetails(c *gin.Context) {
	username := c.Param("username")
	userTable := repository.GetTestUsersTable()
	search_user, err := userTable.GetUser(username)
	if err != nil || search_user.Username == "" {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, search_user)
		return
	}
}
