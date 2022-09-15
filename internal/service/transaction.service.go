package service

import (
	"errors"
	"fmt"
	"roof-stack/internal/dto"
	"roof-stack/internal/entity"
	"roof-stack/internal/repository"
)

type ITransactionService interface {
	Create(dto dto.CreateTransactionDTO) (interface{}, error)
	List( requestParams map[string]interface{}) ([]*entity.Transaction, error)
	Get(id uint64 ) (*entity.Transaction, error)
}


type transactionService struct {
	transactionRepository repository.ITransactionRepository
	walletRepository repository.IWalletRepository
}

func NewTransactionService(
		transactionRepository repository.ITransactionRepository, 
		walletRepository repository.IWalletRepository) *transactionService {
	return &transactionService{
			transactionRepository: transactionRepository,
			walletRepository: walletRepository,
		}
}

func(s *transactionService) Create(dto dto.CreateTransactionDTO) (interface{}, error) {
		fmt.Println(dto.WalletID, dto.UserID)
		wallet, err := s.walletRepository.Get(dto.WalletID, dto.UserID)
		if err != nil {
			return nil, err
		}
		walletUpdate, err := s.walletRepository.Update(dto, wallet)
		if err != nil {
			fmt.Println(walletUpdate)

			return  walletUpdate, err
		}

		transactionCreate, err := s.transactionRepository.Create(dto)
		if err != nil {
			return  transactionCreate, errors.New("Update error")
		}
		return "Transaction created successfully", nil
}

func(s *transactionService) List(requestParams map[string]interface{}) ([]*entity.Transaction, error) {
	res, err := s.transactionRepository.List(requestParams)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func(s *transactionService) Get(id uint64) (*entity.Transaction, error) {
	res, err := s.transactionRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
