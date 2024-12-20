package models

import (
	"database/sql"
)

type User struct {
    ID        int             `json:"id" gorm:"primaryKey"`
    Username  string          `json:"username" gorm:"unique;not null;type:varchar(50)" validate:"required,min=3,max=50"`
    Email     sql.NullString  `json:"email" gorm:"unique;type:varchar(100)" validate:"email"`
    Password  string          `json:"-" gorm:"type:varchar(100)" validate:"required,min=6,max=100"`
    CreatedAt string          `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt string          `json:"updated_at" gorm:"autoUpdateTime"`
	IsStaff   bool            `json:"is_staff" gorm:"default:false"`
}
