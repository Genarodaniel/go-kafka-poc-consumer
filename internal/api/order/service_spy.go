package order

import "context"

type OrderServiceSpy struct {
	OrderServiceInterface
	PostOrderResponse PostOrderResponse
	PostOrderError    error
}

func (s OrderServiceSpy) PostOrder(ctx context.Context, order *PostOrderRequest) (*PostOrderResponse, error) {
	return &s.PostOrderResponse, s.PostOrderError
}
