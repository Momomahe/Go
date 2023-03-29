package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name_book string    `gorm:"not null;type:varchar(100)" json:"name_book"`
	Author    string    `gorm:"not null;type:varchar(100)" json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now
	return nil
}

func (b *Book) BeforeUpdate(tx *gorm.DB) error {
	b.UpdatedAt = time.Now()
	return nil
}
