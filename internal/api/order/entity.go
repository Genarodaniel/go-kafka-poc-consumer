package order

import (
	"errors"

	"github.com/google/uuid"
)

type PostOrderResponse struct {
	OrderID string `json:"order_id"`
}

type PostOrderRequest struct {
	Amount   float64 `json:"amount"`
	ClientID string  `json:"client_id"`
	StoreID  string  `json:"store_id"`
}

func (request *PostOrderRequest) Validate() error {
	if request.Amount <= 0 {
		return errors.New("amount must be a positive number")
	}

	if err := uuid.Validate(request.ClientID); err != nil {
		return errors.New("client_id must be a uuid")
	}

	if err := uuid.Validate(request.StoreID); err != nil {
		return errors.New("store_id must be a uuid")
	}

	return nil
}
