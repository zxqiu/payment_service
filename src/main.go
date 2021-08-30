package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/echo", getServerStatus)

	router.Run("localhost:1111")
}

func getServerStatus(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "running"})
}
