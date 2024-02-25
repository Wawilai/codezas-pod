package entity

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	AccountId uint `gorm:"not null"`
	Status    uint `gorm:"not null"`
}
