package repository

import (
	"errors"
	"fmt"
	"roof-stack/internal/dto"
	"roof-stack/internal/entity"

	"gorm.io/gorm"
)

type ITransactionRepository interface {
	Create(dto dto.CreateTransactionDTO) (entity.Transaction, error)
	List( requestParams map[string]interface{} ) ([]*entity.Transaction, error)
	Get(id uint64) (*entity.Transaction, error)
}

type transactionRepository struct {
	connection *gorm.DB
}

func NewTransactionRepository(dbConn *gorm.DB) *transactionRepository{
	return &transactionRepository{ connection: dbConn}
}

func(db *transactionRepository) Create(dto dto.CreateTransactionDTO) (entity.Transaction, error) {
	transaction := entity.Transaction{}
	transaction.Amount = dto.Amount
	transaction.Status = true
	transaction.WalletID = dto.WalletID
	transaction.CurrencyID = dto.CurrencyID
	transaction.ActionType = dto.ActionType
	result := db.connection.Create(&transaction)
	if result.RowsAffected == 0 {
		return transaction, errors.New("No data found : create")
	}
	return transaction, nil
}

func (db *transactionRepository) List(requestParams map[string]interface{}) ([]*entity.Transaction, error)  {
	var wallets []*entity.Transaction
	fmt.Println(requestParams["limit"])
	result := db.connection.Where("created_at BETWEEN ? AND ?",
									 requestParams["from"], requestParams["to"] ).
							Where("status = ?", requestParams["status"]).
							Limit(requestParams["limit"].(int)).
						Find(&wallets, "status = ? ", "true")
    if result.RowsAffected == 0 {
		return nil, errors.New("No data found")
	}

	return wallets, nil
}

func (db *transactionRepository) Get(id uint64) (*entity.Transaction, error)  {
	var wallet *entity.Transaction
	result := db.connection.Find(&wallet, "status = ? ", "true").Where("id = ?", id)
	if result.RowsAffected == 0 {
		return nil, errors.New("No data found")
	}

	return wallet, nil 
}