package model

type TransactionType string
type TransactionStatus string

const (
	Debit  TransactionType = "DEBIT"
	Credit TransactionType = "CREDIT"
)

const (
	Success TransactionStatus = "SUCCESS"
	Failed  TransactionStatus = "FAILED"
	Pending TransactionStatus = "PENDING"
)

type Transaction struct {
	Timestamp   int64             `json:"timestamp"`
	Name        string            `json:"name"`
	Type        TransactionType   `json:"type"`
	Amount      int64             `json:"amount"`
	Status      TransactionStatus `json:"status"`
	Description string            `json:"description"`
}
