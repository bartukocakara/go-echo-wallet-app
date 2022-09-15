package dto

type CreateWalletDTO struct {
	Title      string  `json:"title"`
	UserID     uint64  `json:"user_id,omitempty"`
	CurrencyID uint64  `json:"currency_id"`
	Balance    float64 `json:"balance,omitempty"`
	Limit      float64 `json:"limit,omitempty"`
	Amount     float64 `json:"amount"`
	Status     bool    `json:"status"`
}

type UpdateWalletDTO struct {
	Title      string  `json:"title"`
	UserID     uint64  `json:"user_id"`
	CurrencyID uint64  `json:"currency_id,omitempty"`
	Balance    float64 `json:"balance,omitempty"`
	Limit      float64 `json:"limit,omitempty"`
	Status     float64 `json:"status,omitempty"`
}