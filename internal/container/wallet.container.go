package container

import (
	handler "roof-stack/internal/handler/wallet"
	"roof-stack/internal/repository"
	"roof-stack/internal/service"
)

var (
	IwalletRepository repository.IWalletRepository = repository.NewWalletRepository(DB)
	IwalletService    service.IWalletService       = service.NewWalletService(IwalletRepository)
	IwalletHandler handler.IWalletHandler = handler.NewWalletHandler(IwalletService)
)