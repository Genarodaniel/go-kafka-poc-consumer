package order

type PostOrderResponse struct {
	OrderID string `json:"order_id"`
}

type PostOrderRequest struct {
	Amount            float64 `json:"amount"`
	OrderID           string  `json:"order_id"`
	ClientID          string  `json:"client_id"`
	StoreID           string  `json:"store_id"`
	NotificationEmail string  `json:"notification_email"`
	Status            string  `json:"status"`
}
