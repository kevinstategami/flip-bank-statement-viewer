package main

import (
	"flip-bank-statement-viewer/internal/handler"
	"flip-bank-statement-viewer/internal/repository"
	"flip-bank-statement-viewer/internal/service"
	"flip-bank-statement-viewer/internal/storage"
	"log"
	"net/http"
)

func main() {
	store := storage.NewMemoryStore()
	repo := repository.NewTransactionRepository(&store.Transactions)
	svc := service.NewTransactionService(repo)
	h := handler.NewTransactionHandler(svc)

	http.HandleFunc("/upload", h.Upload)
	http.HandleFunc("/balance", h.GetBalance)
	http.HandleFunc("/issues", h.GetIssues)

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
