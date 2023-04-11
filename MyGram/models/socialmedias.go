package models

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GORMModel
	Name             string    `gorm:"not null" json:"name" validate:"required-Name is required"`
	Social_Media_Url string    `gorm:"not null" json:"social_media_url" validate:"required-Social_Media_Url is required"`
	UserID           uint      `gorm:"not null" json:"user_id"`
	CreatedAt        time.Time `json:"created_at",omitempty`
	UpdatedAt        time.Time `json:"updated_at",omitempty`
}

func (u *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
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
