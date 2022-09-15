package repository

import (
	"errors"
	"fmt"
	"roof-stack/internal/dto"
	"roof-stack/internal/entity"

	"gorm.io/gorm"
)

type IWalletRepository interface {
	Get(id, userId uint64) (entity.Wallet, error)
	Create(entityWallet entity.Wallet) (entity.Wallet, error)
	Update(dto dto.CreateTransactionDTO, wallet entity.Wallet) (interface{}, error)
	List() []entity.Wallet
}

type walletRepository struct {
	connection *gorm.DB
}

func NewWalletRepository(dbConn *gorm.DB) *walletRepository{
	return &walletRepository{ connection: dbConn}
}

func (db *walletRepository) Get(id, userId uint64) (entity.Wallet, error) {
	var wallet entity.Wallet
    db.connection.
				  Where("user_id = ?", userId).
				  Where("id = ?", id).
				  Where("status", true).
				  Find(&wallet)
    return wallet, nil
}

func(db *walletRepository) Create(entityWallet entity.Wallet) (entity.Wallet, error) {
	db.connection.Save(&entityWallet)
	result := db.connection.Joins("User", "Currency").Find(&entityWallet)
	if result.RowsAffected == 0 {
		return entityWallet, errors.New("No store found")
	}
	return entityWallet, nil
}

func(db *walletRepository) Update(dto dto.CreateTransactionDTO, wallet entity.Wallet) (interface{}, error) {
	actionType := dto.ActionType
	amount := 0.00
	switch actionType {
		case entity.WITHDRAW:
			amount = wallet.Balance - dto.Amount
			fmt.Println(actionType)
			if wallet.Balance < dto.Amount {
				return nil, errors.New("Not enough balance")
			}

		case entity.DEPOSIT:
			amount =wallet.Balance + dto.Amount
		}
	result := db.connection.Model(&wallet).
						Where("id", dto.WalletID).
						Where("user_id", dto.UserID).
						Update("balance", amount)
	if result.RowsAffected == 0 {
		return nil, errors.New("No wallet found")
	}
	return result, nil
}

func updateCase() {

}

func (db *walletRepository) List() []entity.Wallet {
	var wallets []entity.Wallet  
    db.connection.Find(&wallets)

    return wallets 
}

