package dto

import "roof-stack/internal/entity"

type TransactionType string

type CreateTransactionDTO struct {
	UserID     uint64            `json:"user_id,omitempty"`
	WalletID   uint64            `json:"wallet_id,,omitempty"`
	CurrencyID uint64            `json:"currency_id,,omitempty"`
	ActionType entity.ActionType `json:"action_type,omitempty"`
	Amount     float64           `json:"amount"`
}

type UpdateTransactionDTO struct {
	Title  string `json:"title"`
	Status bool   `json:"status"`
}