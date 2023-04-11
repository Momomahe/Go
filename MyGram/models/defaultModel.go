package models

type GORMModel struct {
	ID uint `gorm:"primaryKey" json:"id"`
}
