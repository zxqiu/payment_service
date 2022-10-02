package actions

import (
	"log"
)

type CreatePaymentMetadata struct {
	ActionID string
}

func (metadata CreatePaymentMetadata) GetActionID() string {
	return metadata.ActionID
}

type CreatePaymentAction struct{}

func (action CreatePaymentAction) Execute(metadata ActionMetadata) error {
	log.Println("Creating payment")
	log.Printf("Action ID : %s\n", metadata.GetActionID())

	return nil
}
