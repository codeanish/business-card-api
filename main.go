package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(corsMiddleware())
	router.GET("/", healthCheck)
	router.GET("/users/:username", getUserDetails)
	router.GET("/users/:username/leetcode", getLeetcodeRanking)
	router.GET("/users/:username/github", getGithubTotalCommits)
	router.GET("/users/:username/stackoverflow/:userId", getStackOverflowReputation)
	router.Run()
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
