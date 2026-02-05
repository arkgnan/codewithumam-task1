package services

import (
	"crud-category/models"
	"crud-category/repositories"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{
		repo: repo,
	}
}

func (s *TransactionService) Checkout(items []models.CheckoutItem) (*models.Transaction, error) {
	return s.repo.CreateTransaction(items)
}

func (s *TransactionService) TransactionReport(start, end string) (*models.TransactionReport, error) {
	return s.repo.TransactionReport(start, end)
}
