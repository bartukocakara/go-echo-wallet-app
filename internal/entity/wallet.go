package entity

import "time"

type Wallet struct {
	ID               uint64     `gorm:"primary_key:auto_increment" json:"id"`
	UserID           uint64     `gorm:"not null" json:"user_id"`
	User             User       `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	Status           bool     	`gorm:"not null" json:"status"`
	CurrencyID       uint64     `gorm:"not null" json:"currency_id"`
	Currency         Currency   `gorm:"foreignkey:CurrencyID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	Title            string     `gorm:"not null;type:varchar(255)" json:"title"`
	Balance          float64    `gorm:"not null;type:decimal(15,2)" json:"balance"`
	Limit 			 float64    `gorm:"not null;type:decimal(15,2)" json:"limit"`
	CreatedAt        time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}