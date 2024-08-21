package models

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	CompID  uint64  `gorm:"column:compId"`
	EventID uint64  `gorm:"column:eventId"`
	Sold    int     `gorm:"column:sold"`
	Price   float64 `gorm:"column:price"`
	Status  int     `gorm:"column:status"`
}
