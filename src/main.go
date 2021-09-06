package main

import (
	"log"
	"net/http"

	"pay_gate/services"
	"pay_gate/services/actions"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Server started")

	router := gin.Default()
	router.GET("/echo", getServerStatus)

	router.POST("/payments", services.GetAction(actions.CreatePaymentActionType))

	router.Run("localhost:1111")
}

func getServerStatus(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "running"})
}
