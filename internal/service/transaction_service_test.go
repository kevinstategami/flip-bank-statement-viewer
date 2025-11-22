package service

import (
	"flip-bank-statement-viewer/internal/model"
	"flip-bank-statement-viewer/internal/repository"
	"strings"
	"testing"
)

func TestService_GetBalance(t *testing.T) {
	data := []model.Transaction{
		{Type: model.Credit, Amount: 1000, Status: model.Success},
		{Type: model.Debit, Amount: 400, Status: model.Success},
		{Type: model.Debit, Amount: 50, Status: model.Pending},
	}

	repo := repository.NewTransactionRepository(&data)
	svc := NewTransactionService(repo)

	balance := svc.GetBalance()
	expected := int64(600)

	if balance != expected {
		t.Fatalf("expected %d got %d", expected, balance)
	}
}

func TestService_GetIssues(t *testing.T) {
	data := []model.Transaction{
		{Status: model.Success},
		{Status: model.Pending},
		{Status: model.Failed},
	}

	repo := repository.NewTransactionRepository(&data)
	svc := NewTransactionService(repo)

	issues := svc.GetIssues()
	if len(issues) != 2 {
		t.Fatalf("expected 2 issues, got %d", len(issues))
	}
}

func TestService_Upload_Invalid(t *testing.T) {
	data := []model.Transaction{
		{Type: "WRONG", Status: model.Success, Amount: 100},
	}

	repo := repository.NewTransactionRepository(&[]model.Transaction{})
	svc := NewTransactionService(repo)

	err := svc.Upload(data)
	if err == nil {
		t.Fatalf("expected validation error, got nil")
	}

	if !strings.Contains(err.Error(), "invalid transaction type at row") {
		t.Fatalf("unexpected error message: %v", err)
	}
}

func TestService_Upload_Success(t *testing.T) {
	data := []model.Transaction{
		{Type: model.Credit, Status: model.Success, Amount: 100},
	}

	storage := []model.Transaction{}
	repo := repository.NewTransactionRepository(&storage)
	svc := NewTransactionService(repo)

	err := svc.Upload(data)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(storage) != 1 {
		t.Fatalf("expected 1 stored row, got %d", len(storage))
	}
}
