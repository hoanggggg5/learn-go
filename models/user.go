package models

import (
	"time"
)

type User struct {
	ID        int64     `gorm:"primaryKey"`
	UID       string    `gorm:"type:character varying(10);not null"`
	Email     string    `gorm:"type:character varying(50);not null"`
	Password  string    `gorm:"type:text;not null"`
	Role      string    `gorm:"type:character varying(10);not null;default:member"`
	Status    string    `gorm:"type:character varying(10);not null;default:active"`
	CreatedAt time.Time `gorm:"type:timestamp(0)"`
	UpdatedAt time.Time `gorm:"type:timestamp(0)"`
}
