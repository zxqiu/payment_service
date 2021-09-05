package actions

import (
	"log"
)

type CreatePaymentAction struct {
}

func (action CreatePaymentAction) Execute(metadata ActionMetadata) error {
	log.Println("Creating payment")

	return nil
}
