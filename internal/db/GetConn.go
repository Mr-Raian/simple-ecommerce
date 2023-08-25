package db

import "github.com/jmoiron/sqlx"

func (rcv postgres) GetConn() *sqlx.DB {
	return rcv.DB
}
