package order_test

import (
	"go-kafka-order-producer/internal/api/order"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("Should return an error when amount is empty", func(t *testing.T) {
		request := order.PostOrderRequest{}
		err := request.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "amount must be a positive number")
	})

	t.Run("Should return an error when amount is negative", func(t *testing.T) {
		request := order.PostOrderRequest{
			Amount: -123.00,
		}
		err := request.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "amount must be a positive number")
	})

	t.Run("Should return an error when clientID is empty", func(t *testing.T) {
		request := order.PostOrderRequest{
			Amount: 123.00,
		}
		err := request.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "client_id must be a uuid")
	})

	t.Run("Should return an error when clientID is not a uuid", func(t *testing.T) {
		request := order.PostOrderRequest{
			Amount:   123.00,
			ClientID: "not a uuid",
		}
		err := request.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "client_id must be a uuid")
	})

	t.Run("Should return an error when storeID is empty", func(t *testing.T) {
		request := order.PostOrderRequest{
			Amount:   123.00,
			ClientID: uuid.NewString(),
		}
		err := request.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "store_id must be a uuid")
	})

	t.Run("Should return an error when storeID is not a uuid", func(t *testing.T) {
		request := order.PostOrderRequest{
			Amount:   123.00,
			ClientID: uuid.NewString(),
			StoreID:  "not a uuid",
		}
		err := request.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "store_id must be a uuid")
	})

	t.Run("Should return success when the request is valid", func(t *testing.T) {
		request := order.PostOrderRequest{
			Amount:   123.00,
			ClientID: uuid.NewString(),
			StoreID:  uuid.NewString(),
		}
		err := request.Validate()

		assert.Nil(t, err)
	})

}
