package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ID          string
	Name        string
	Price       string
	Description string
}
