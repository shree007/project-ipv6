package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	NotFoundError = errors.New("Path Incorrect, please use '/' path to access viz. /?token=SOMETOKENVALUE")
)

type applicationOneResponse struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

const fixedToken = "applicationtwo_secure_token"

func main() {
	log.Info("creating gin router")
	router := SetupRouter()
	router.Run(":8090")
}

func SetupRouter() *gin.Engine {
	log.Info("Enter into SetupRouter Function")
	serviceEndpoint := getServiceEndpoint()
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

		response, err := http.Get(serviceEndpoint + "/?token=bXlfc2VjdXJlX3Rva2Vu")

		if err != nil {
			log.Printf("Failed to reach service %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Application one is unreachable, please check"})
			return
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)

		var applicationone applicationOneResponse

		if err := json.Unmarshal(body, &applicationone); err != nil {
			c.Error(errors.New("failed to parse response which we recieved from application one"))
			return
		}
		reverseMsg := reverseString(applicationone.Message)

		c.JSON(http.StatusOK, gin.H{
			"id":      "1",
			"message": reverseMsg,
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

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func getServiceEndpoint() string {
	appEnv := os.Getenv("APP_ENV")
	switch appEnv {
	case "docker":
		return "http://application01:8080"
	case "kubernetes":
		return "http://application-01-service.project-ipv6.svc.cluster.local"
	default:
		return "http://localhost:8080"
	}
}
