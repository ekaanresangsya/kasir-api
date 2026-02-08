package service

import (
	"kasir-api/internal/model"
	"kasir-api/internal/repository"
)

type TransactionService struct {
	transactionRepo *repository.TransactionRepository
}

func NewTransactionService(transactionRepo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{transactionRepo: transactionRepo}
}

func (s *TransactionService) Checkout(req model.CheckoutRequest) (*model.Transaction, error) {
	return s.transactionRepo.CreateTransaction(req)
}
