package db

import (
	"api/internal/data"
	"context"
	"fmt"
	"strings"
)

func (rcv postgres) GetItemByID(ctx context.Context, id string, columns ...string) (data.Item, error) {
	p := data.Item{}
	columnsToSelect := "id, price, metadata"
	if len(columns) != 0 {
		columnsToSelect = strings.Join(columns, ", ")
	}
	if err := rcv.DB.GetContext(ctx, &p, fmt.Sprintf("SELECT %s FROM items WHERE id = $1 LIMIT 1;", columnsToSelect), id); err != nil {
		return p, err
	}

	return p, nil
}
