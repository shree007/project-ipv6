package main

import (
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	NotFoundError = errors.New("Path Incorrect, please use '/' path to access")
)

const fixedToken = "my_secure_token"

func main() {
	log.Info("creating gin router")
	router := SetupRouter()
	router.Run(":8080")
}

func SetupRouter() *gin.Engine {
	log.Info("Enter into SetupRouter Function")
	router := gin.Default()

	router.Use(ErrorHandling())

	router.GET("/", func(c *gin.Context) {
		encoded_token := c.Query("token")
		decoded_token, err := base64.StdEncoding.DecodeString(encoded_token)
		if err != nil {
			log.Errorf("failed to decode Token %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Base64 Token"})
			return
		}
		if string(decoded_token) != fixedToken {
			log.Warn("Invalid Token")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized Invalid Token, Please use correct One"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":      "1",
			"message": "Hello world",
		})
	})

	router.NoRoute(func(c *gin.Context) {
		customizeErrorHandling(c, NotFoundError)
	})
	return router
}

func ErrorHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			customizeErrorHandling(c, err)
		}
	}
}

func customizeErrorHandling(c *gin.Context, err error) {
	log.Info("Enter into customizeErrorHandling Function")
	switch err {
	case NotFoundError:
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error, Please check resources"})
	}
}
