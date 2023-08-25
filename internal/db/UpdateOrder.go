package db

import (
	"api/internal/data"
	"context"
	"errors"
	"fmt"
)

func (rcv postgres) UpdateOrder(ctx context.Context, order data.Order, where data.Where, columns ...string) error {
	query := "UPDATE orders SET "
	for i, column := range columns {
		if i > 0 {
			query += ", "
		}
		query += column + " = :" + column
	}
	if where == (data.Where{}) {
		return errors.New("empty Where")
	}
	query += " WHERE " + where.Column + " = '" + where.EqualsTo + "'"

	// Execute the update query
	fmt.Println(query)
	_, err := rcv.DB.NamedExecContext(ctx, query, order)
	if err != nil {
		return err
	}

	return nil
}
