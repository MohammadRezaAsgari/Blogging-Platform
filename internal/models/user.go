package models

import "time"

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null;type:varchar(50)"`
	Password string `json:"-" gorm:"type:varchar(100)"`
	JoinedAt time.Time `json:"joined_at" gorm:"autoCreateTime"`
	IsStaff  bool   `json:"is_staff" gorm:"default:false"`
}
