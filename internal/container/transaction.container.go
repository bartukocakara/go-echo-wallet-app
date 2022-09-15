package container

import (
	handler "roof-stack/internal/handler/transaction"
	"roof-stack/internal/repository"
	"roof-stack/internal/service"
)

var (
	ItransactionRepository repository.ITransactionRepository = repository.NewTransactionRepository(DB)
	ItransactionService    service.ITransactionService       = service.NewTransactionService(ItransactionRepository, IwalletRepository)
	ItransactionHandler handler.ITransactionHandler = handler.NewTransactionHandler(ItransactionService)
)