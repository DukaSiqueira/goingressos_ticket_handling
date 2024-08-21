package models

import (
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	CompID  uint64
	EventID uint64
	Name    string
	Sold    int
	Price   float64
	Status  int
}
