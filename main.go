package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func add(a, b int) int {
	return a + b
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
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
			"result": add(a, b),
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
