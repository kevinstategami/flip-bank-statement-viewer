package repository

import "flip-bank-statement-viewer/internal/model"

type TransactionRepository interface {
	SaveAll([]model.Transaction)
	FindAll() []model.Transaction
}

type transactionRepository struct {
	storage *[]model.Transaction
}

func NewTransactionRepository(storage *[]model.Transaction) TransactionRepository {
	return &transactionRepository{
		storage: storage,
	}
}

func (r *transactionRepository) SaveAll(data []model.Transaction) {
	*r.storage = make([]model.Transaction, len(data))
	copy(*r.storage, data)
}

func (r *transactionRepository) FindAll() []model.Transaction {
	return *r.storage
}
