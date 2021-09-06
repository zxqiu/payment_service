package events

type EventType string

const (
	CreatePaymentStartEventType     EventType = "CreatePaymentStartEventType"
	CreatePaymentSucceededEventType EventType = "CreatePaymentSucceededEventType"
	CreatePaymentFailedEventType    EventType = "CreatePaymentFailedEventType"
	AuthStartEventType              EventType = "AuthStartEventType"
	AuthSucceededEventType          EventType = "AuthSucceededEventType"
	AuthFailedEventType             EventType = "AuthFailedEventType"
)
