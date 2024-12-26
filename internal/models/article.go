package models

import "time"

type Article struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null;type:varchar(50)"`
	Body      string    `json:"body" gorm:"not null;type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	UserID    int       `json:"-" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
}
