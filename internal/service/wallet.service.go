package service

import (
	"roof-stack/internal/dto"
	"roof-stack/internal/entity"
	"roof-stack/internal/repository"

	"github.com/mashingan/smapping"
	log "github.com/sirupsen/logrus"
)

type IWalletService interface {
	Create(dto dto.CreateWalletDTO) (entity.Wallet, error)
	List( ) ([]entity.Wallet, error)
}


type walletService struct {
	WalletRepository repository.IWalletRepository
}

func NewWalletService(WalletRepository repository.IWalletRepository) *walletService {
	return &walletService{WalletRepository: WalletRepository}
}

func(s *walletService) Create(dto dto.CreateWalletDTO) (entity.Wallet, error) {
	wallet := entity.Wallet{}
	dto.Status = true
	err := smapping.FillStruct(&wallet, smapping.MapFields(dto))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res, err := s.WalletRepository.Create(wallet)
	return res, nil
}

func(s *walletService) List() ([]entity.Wallet, error) {
	res := s.WalletRepository.List()
	return res, nil
}
