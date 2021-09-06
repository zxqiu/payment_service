package events

import (
	"log"
	"pay_gate/services/actions"
)

type EventMetadata interface{}

func TriggerAction(eventType EventType, matadata *EventMetadata) error {
	log.Println("Triggering event", event)
	var action actions.Action
	switch eventType {
	case CreatePaymentStartEventType:
		action = actions.CreatePaymentAction{}
	case CreatePaymentSucceededEventType:
		// call func
	case CreatePaymentFailedEventType:
		// call func
	default:
		return error{}
	}

	if action != nil {
		err := action.Execute(nil)
		return err
	}

	return nil
}
