package main

import (
	"net/http"

	"pay_gate/services"
	"pay_gate/services/actions"
	"pay_gate/sm"

	"github.com/gin-gonic/gin"
)

func main() {
	sm.TriggerEvent(sm.ServerStartedEvent, nil)

	router := gin.Default()
	router.GET("/echo", getServerStatus)

	router.POST("/payments", services.GetAction(actions.CreatePaymentActionType))

	router.Run("localhost:1111")
}

func getServerStatus(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "running"})
}
