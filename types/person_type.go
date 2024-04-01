package types

import "gorm.io/gorm"

type PersonType struct {
	gorm.Model
	Username string `gorm:"not null;unique" json:"username" binding:"required"`
	Email    string `gorm:"not null;unique" json:"email" validate:"required,email"`
	Phone    string `gorm:"not null;unique" json:"phone" binding:"required,min=10,max=13"`
	Name     string `gorm:"not null" json:"name" binding:"required"`
	Address  string `gorm:"not null" json:"address"`
	Age      uint   `json:"age"`
	Gender   string `json:"gender"`
}
