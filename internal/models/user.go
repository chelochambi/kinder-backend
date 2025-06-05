package models

import (
	"time"
)

type User struct {
	ID           uint   `gorm:"primary_key"`
	Username     string `gorm:"type:varchar(100);unique_index"`
	Email        string `gorm:"type:varchar(100);unique_index"`
	PasswordHash string `gorm:"not null"`
	FirstName    string `gorm:"type:varchar(100)"`
	LastName     string `gorm:"type:varchar(100)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (user *User) TableName() string {
	return "usuarios" // Si tienes un nombre de tabla específico
}
