package models

import "time"

// Ticket representa a tabela de ingressos no banco de dados
type Ticket struct {
	ID        uint   `gorm:"primaryKey"`
	EventID   uint   `gorm:"column:eventId;not null"` // Mapeia explicitamente a coluna para eventId
	UserID    uint   `gorm:"not null"`
	Status    string `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
