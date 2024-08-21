package models

import "gorm.io/gorm"

type Buy struct {
	gorm.Model
	CompID     uint64  `gorm:"column:compId"`
	EventID    uint64  `gorm:"column:eventId"`
	CustomerID *uint64 `gorm:"column:customerId"`
	BatchID    uint64  `gorm:"column:batchId"`
	Status     int
	Price      float64
	Quantity   int
	Hash       string
}
