package service

import (
	"errors"

	"flip-bank-statement-viewer/internal/model"
	"flip-bank-statement-viewer/internal/repository"
)

type TransactionService interface {
	Upload([]model.Transaction) error
	GetBalance() int64
	GetIssues() []model.Transaction
}

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	return &transactionService{repo: repo}
}

func (s *transactionService) Upload(data []model.Transaction) error {
	if len(data) == 0 {
		return errors.New("no transactions found")
	}

	for i, t := range data {
		if !isValidType(t.Type) {
			return errors.New("invalid transaction type at row " + string(i) + ": " + string(t.Type))
		}
		if !isValidStatus(t.Status) {
			return errors.New("invalid transaction status at row " + string(i) + ": " + string(t.Status))
		}
		if t.Amount < 0 {
			return errors.New("invalid amount: cannot be negative")
		}
	}

	s.repo.SaveAll(data)
	return nil
}

func (s *transactionService) GetBalance() int64 {
	var balance int64 = 0

	for _, t := range s.repo.FindAll() {
		if t.Status != model.Success {
			continue
		}
		switch t.Type {
		case model.Credit:
			balance += t.Amount
		case model.Debit:
			balance -= t.Amount
		}
	}

	return balance
}

func (s *transactionService) GetIssues() []model.Transaction {
	var issues []model.Transaction
	for _, t := range s.repo.FindAll() {
		if t.Status == model.Failed || t.Status == model.Pending {
			issues = append(issues, t)
		}
	}
	return issues
}

func isValidType(t model.TransactionType) bool {
	return t == model.Debit || t == model.Credit
}

func isValidStatus(s model.TransactionStatus) bool {
	return s == model.Success || s == model.Failed || s == model.Pending
}
