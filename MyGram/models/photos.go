package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type Photo struct {
	GORMModel
	Title     string `gorm:"not null" json:"title" validate:"required-Title is required"`
	Caption   string `gorm:"not null" json:"caption" validate:"required-Caption is required"`
	Photo_Url string `gorm:"not null" json:"photo_url" validate:"required-Photo_Url is required"`
	UserID    uint   `gorm:"not null" json:"user_id"`
}

func (u *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	userData, exists := tx.Statement.Context.Value("userData").(jwt.MapClaims)
	if !exists {
		return errors.New("userData is missing in the context")
	}

	u.UserID = uint(userData["id"].(float64))

	_, err = govalidator.ValidateStruct(u)

	if err != nil {
		return
	}

	return
}
