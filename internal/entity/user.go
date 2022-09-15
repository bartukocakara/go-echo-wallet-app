package entity

import "time"

type User struct {
	ID        uint64     `gorm:"primary_key:auto_increment" json:"id"`
	FirstName string     `gorm:"type:varchar(255)" json:"first_name" faker:"first_name"`
	LastName  string     `gorm:"type:varchar(255)" json:"last_name"`
	Email     string     `gorm:"uniqueIndex;type:varchar(255)" json:"email" faker:"email"`
	Password  string     `gorm:"->;<-;not null" json:"-" faker:"password"`
	Token     string     `gorm:"-" json:"token,omitempty"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
