package models

import (
	"mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GORMModel
	UserName string `gorm:"not null" json:"username" validate:"required-First name is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" validate:"required-Email is required,email-Invalid email format"`
	Password string `gorm:"not null" json:"password" validate:"required-Password is required,MinStringLength(6)-Password has to have a minimum length of 6 characters"`
	Age      uint8  `gorm:"not null" json:"age" validate:"required,MinUint8(8)-Age must be greater than or equal to 8"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return
	}

	hashedPass, err := helpers.HashPassword(u.Password)
	if err != nil {
		return
	}
	u.Password = hashedPass

	return
}
