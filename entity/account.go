package entity

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Account_Id uint `gorm:"not null"`
	Status     uint `gorm:"not null"`
}
