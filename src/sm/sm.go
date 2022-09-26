package sm

import (
	"log"
	"pay_gate/services/actions"
)

type EventType string

type EventMetadata interface{}

type EventHandler interface {
	Execute(metadata EventMetadata)
}

const (
	ServerStartedEvent          EventType = "ServerStartedEvent"
	CreatePaymentStartEvent     EventType = "CreatePaymentStartEvent"
	CreatePaymentSucceededEvent EventType = "CreatePaymentSucceededEvent"
	CreatePaymentFailedEvent    EventType = "CreatePaymentFailedEvent"
	AuthStartEvent              EventType = "AuthStartEvent"
	AuthSucceededEvent          EventType = "AuthSucceededEvent"
	AuthFailedEvent             EventType = "AuthFailedEvent"
)

func TriggerEvent(event EventType, matadata *EventMetadata) error {
	log.Println("Triggering event", event)
	var action actions.Action
	switch event {
	case ServerStartedEvent:
		log.Println("Server started")
	case CreatePaymentStartEvent:
		action = actions.CreatePaymentAction{}
	case CreatePaymentSucceededEvent:
		// call func
	case CreatePaymentFailedEvent:
		// call func
	default:
		return nil
	}

	if action != nil {
		return action.Execute(nil)
	}

	return nil
}
