package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/ZBIP/ZBIP/pkg/add"
	"github.com/gin-gonic/gin"
)

type Secret struct {
	Key   int    `json:"number"`
	Value string `json:"value"`
}

type Config struct {
	Secrets []Secret `json:"secretValues"`
}

func main() {
	data := os.Getenv("APP_CONFIG")
	res, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(fmt.Errorf("Failed to decode APP_CONFIG: %w", err))
	}
	cfg := Config{}
	if err := json.Unmarshal(res, &cfg); err != nil {
		panic(fmt.Errorf("Failed to unmarshal APP_CONFIG: %w", err))
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Vulnerable code
	r.GET("/download", func(c *gin.Context) {
		dir := "/Users/{CHANGE_PROJECT_DIRECTRY}/"

		// Although the file name is hard-coded, we assume that the file name is actually determined by the DB or user input.
		filename := "malicious.sh\";dummy=.txt"
		c.FileAttachment(dir+filename, filename)
	})

	r.GET("/add/:a/:b", func(c *gin.Context) {
		a, err := strconv.Atoi(c.Param("a"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid parameter a",
			})
		}
		b, err := strconv.Atoi(c.Param("b"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid parameter b",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"result": add.Add(a, b),
		})
	})

	r.GET("/sub/:a/:b", func(c *gin.Context) {
		a, err := strconv.Atoi(c.Param("a"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid parameter a",
			})
		}
		b, err := strconv.Atoi(c.Param("b"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid parameter b",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"result": a-b,
		})
	})

	r.GET("/getsecret/:a", func(c *gin.Context) {
		a, err := strconv.Atoi(c.Param("a"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid parameter a",
			})
			return
		}

		for _, secret := range cfg.Secrets {
			if secret.Key == a {
				c.JSON(http.StatusOK, gin.H{
					"result": secret.Value,
				})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Secret not found",
		})
	})

	if err := r.Run(); err != nil {
		panic(err)
	}
}
