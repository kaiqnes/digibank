package entities

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	AccountID       uint // FK
	OperationTypeID uint // FK
	Amount          float64
	EventDate       time.Time // AutoDate
}
