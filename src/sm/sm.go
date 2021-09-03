package sm

type EventType string

type EventMetadata interface{}

type EventHandler interface {
	Execute(metadata EventMetadata)
}

const (
	CreatePayment          EventType = "CreatePayment"
	CreatePaymentSucceeded EventType = "CreatePaymentSucceeded"
	CreatePaymentFailed    EventType = "CreatePaymentFailed"
	AuthStart              EventType = "AuthStart"
	AuthSucceeded          EventType = "AuthSucceeded"
	AuthFailed             EventType = "AuthFailed"
)

func TriggerEvent(event EventType, matadata EventMetadata) {
	switch event {
	case CreatePayment:
		// call func
	case CreatePaymentSucceeded:
		// call func
	case CreatePaymentFailed:
		// call func
	}
}
