package db

import (
	"api/internal/data"
	"context"
	"fmt"
	"strings"
)

func (rcv postgres) GetOrderByID(ctx context.Context, id string, columns ...string) (data.Order, error) {
	o := data.Order{}
	columnsToSelect := "id, item_id, price, email, payment_method, payment_url, payment_id, order_status"
	if len(columns) != 0 {
		columnsToSelect = strings.Join(columns, ", ")
	}
	if err := rcv.DB.GetContext(ctx, &o, fmt.Sprintf("SELECT %s FROM orders WHERE id = $1 LIMIT 1;", columnsToSelect), id); err != nil {
		return o, err
	}

	return o, nil
}
