package data

import "github.com/google/uuid"

type Item struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Price    uint      `json:"price" db:"price"`
	Metadata string    `json:"metadata" db:"metadata"`
}
