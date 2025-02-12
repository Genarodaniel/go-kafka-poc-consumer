package order

import (
	"context"
	"database/sql"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

func (r *OrderRepository) SaveOrder(ctx context.Context, order Order) (string, error) {
	var id string
	var query = `
		INSERT INTO tab_order(
			id,
			store_id,
			client_id,
			current_status,
			notification_email
		)
		VALUES(
			$1,
			$2,
			$3,
			$4,
			$5
		)
		RETURNING id
	`
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return id, err
	}

	defer stmt.Close()

	result, err := stmt.QueryContext(ctx,
		order.ID,
		order.StoreID,
		order.ClientID,
		order.Status,
		order.NotificationEmail,
	)
	if err != nil {
		return id, err
	}

	result.Next()
	err = result.Scan(
		&id,
	)

	return id, nil
}
