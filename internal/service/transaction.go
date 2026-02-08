package service

import (
	"kasir-api/internal/model"
	"kasir-api/internal/repository"
	"time"
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

func (s *TransactionService) GetReportToday() (*model.ReportRes, error) {
	resData := model.ReportRes{}

	// get total transaction today
	now := time.Now()
	today := now.Format("2006-01-02")
	totalTransaction, totalRevenue, err := s.transactionRepo.GetTotalTransaction(today, today)
	if err != nil {
		return nil, err
	}
	resData.TotalTransaksi = totalTransaction
	resData.TotalRevenue = totalRevenue

	// get produk terlaris
	productTerlaris, err := s.transactionRepo.GetProductTerlaris(today, today)
	if err != nil {
		return nil, err
	}
	resData.ProdukTerlaris = *productTerlaris

	return &resData, nil
}

func (s *TransactionService) GetReport(startDate, endDate string) (*model.ReportRes, error) {
	resData := model.ReportRes{}

	today := time.Now().Format("2006-01-02")
	if startDate == "" {
		startDate = today
	}
	if endDate == "" {
		endDate = today
	}

	// get total transaction
	totalTransaction, totalRevenue, err := s.transactionRepo.GetTotalTransaction(startDate, endDate)
	if err != nil {
		return nil, err
	}
	resData.TotalTransaksi = totalTransaction
	resData.TotalRevenue = totalRevenue

	// get produk terlaris
	productTerlaris, err := s.transactionRepo.GetProductTerlaris(startDate, endDate)
	if err != nil {
		return nil, err
	}
	if productTerlaris != nil {
		resData.ProdukTerlaris = *productTerlaris
	}

	return &resData, nil
}
