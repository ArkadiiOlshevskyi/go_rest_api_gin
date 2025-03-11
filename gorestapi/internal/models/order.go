package models

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type Order struct {
	ID        int
	UserID    int
	Items     []Item
	Total     float64
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
