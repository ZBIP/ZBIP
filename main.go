package main

import (
	"net/http"
	"strconv"

	"github.com/ZBIP/ZBIP/pkg/add"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Vulnerable code
	// r.GET("/download", func(c *gin.Context) {
	// 	dir := "/Users/{CHANGE_PROJECT_DIRECTRY}/"

	// 	// Although the file name is hard-coded, we assume that the file name is actually determined by the DB or user input.
	// 	filename := "malicious.sh\";dummy=.txt"
	// 	c.FileAttachment(dir+filename, filename)
	// })

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

	if err := r.Run(); err != nil {
		panic(err)
	}
}
