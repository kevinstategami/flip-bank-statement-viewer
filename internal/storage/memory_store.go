package storage

import "flip-bank-statement-viewer/internal/model"

type MemoryStore struct {
	Transactions []model.Transaction
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		Transactions: make([]model.Transaction, 0),
	}
}
