package main

import (
	"businesscardapi/models"
	"businesscardapi/repository"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

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

func getLeetcodeRanking(c *gin.Context) {
	url := "https://leetcode.com/graphql/"
	method := "POST"
	username := c.Param("username")
	payload := strings.NewReader("{\"query\":\"query userPublicProfile($username: String!) {matchedUser(username: $username) {profile {ranking}}}\",\"variables\":{\"username\":\"" + username + "\"}}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	var leetcodeStats models.LeetcodeStats
	json.Unmarshal([]byte(string(body)), &leetcodeStats)
	c.IndentedJSON(http.StatusOK, gin.H{"ranking": leetcodeStats.Data.MatchedUser.Profile.Ranking})
}

func getGithubTotalCommits(c *gin.Context) {

	method := "GET"
	username := c.Param("username")
	url := fmt.Sprintf("https://api.github.com/search/commits?q=author:%s", username)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Accept", "application/vnd.github.cloak-preview")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	var commitSearchResponse models.CommitSearchResponse
	json.Unmarshal([]byte(string(body)), &commitSearchResponse)
	c.IndentedJSON(http.StatusOK, gin.H{"total_commits": commitSearchResponse.TotalCount})
}
