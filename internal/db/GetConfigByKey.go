package db

import "context"

func (rcv postgres) GetConfigByKey(ctx context.Context, key string) (string, error) {
	var value string
	if err := rcv.DB.GetContext(ctx, &value, "SELECT value FROM config WHERE key = $1 LIMIT 1", key); err != nil {
		return "", err
	}
	return value, nil
}
