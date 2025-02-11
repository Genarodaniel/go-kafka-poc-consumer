package events

type EventType string

const (
	CreateOrder EventType = "order.create"
	UpdateOrder EventType = "order.update"
)
