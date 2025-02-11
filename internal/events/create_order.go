package events

import (
	"time"
)

type CreateOrderEvent struct {
	DateTime time.Time
	Payload  []byte
}

func (c *CreateOrderEvent) GetName() string {
	return "order.create"
}

func (c *CreateOrderEvent) GetDateTime() time.Time {
	return c.DateTime
}

func (c *CreateOrderEvent) GetPayload() []byte {
	return c.Payload
}
