package entity

import "time"

type ActionType string

const (
    WITHDRAW  	ActionType = "WITHDRAW"
    DEPOSIT 	ActionType = "DEPOSIT"
)
type Transaction struct {
	ID                  uint64            `gorm:"primary_key:auto_increment" json:"id"`
	Status   			bool 			  `gorm:"not null" json:"status"`
	ActionType 			ActionType 		  `gorm:"not null" json:"type"`	
	CurrencyID          uint64            `gorm:"not null" json:"currency_type"`
	Currency            Currency          `gorm:"foreignkey:CurrencyID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	WalletID            uint64            `gorm:"not null" json:"to_wallet_id"`
	Wallet           	Wallet            `gorm:"foreignkey:WalletID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	Amount              float64           `gorm:"not null;type:decimal(15,2)" json:"list_price"`
	CreatedAt           time.Time         `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt           time.Time         `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt           *time.Time        `json:"deleted_at,omitempty"`
}