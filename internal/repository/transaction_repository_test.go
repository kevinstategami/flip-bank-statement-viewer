package repository

import (
	"flip-bank-statement-viewer/internal/model"
	"testing"
)

func TestRepository_SaveAndFind(t *testing.T) {
	storage := []model.Transaction{}
	repo := NewTransactionRepository(&storage)

	data := []model.Transaction{
		{Name: "A"},
		{Name: "B"},
	}

	repo.SaveAll(data)
	res := repo.FindAll()

	if len(res) != 2 {
		t.Fatalf("expected 2 results, got %d", len(res))
	}
}
