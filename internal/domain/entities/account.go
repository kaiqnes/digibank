package entities

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	DocumentNumber string
}
