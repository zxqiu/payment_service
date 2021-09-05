package services

import (
	"log"
	"pay_gate/services/actions"

	"github.com/gin-gonic/gin"
)

func GetAction(actionType actions.ActionType) func(c *gin.Context) {
	switch actionType {
	case actions.CreatePaymentActionType:
		return func(c *gin.Context) {
			err := actions.CreatePaymentAction{}.Execute(nil)
			if err != nil {
				log.Fatalln("failed to create payment")
			}
		}
	default:
		return func(c *gin.Context) {}
	}

	return func(c *gin.Context) {}
}
