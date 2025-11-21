package handler

import (
	"encoding/json"
	"flip-bank-statement-viewer/internal/service"
	"flip-bank-statement-viewer/internal/utils"
	"net/http"
	"path/filepath"
	"strings"
)

type TransactionHandler struct {
	svc service.TransactionService
}

func NewTransactionHandler(svc service.TransactionService) *TransactionHandler {
	return &TransactionHandler{svc: svc}
}

func (h *TransactionHandler) Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		http.Error(w, "failed to parse form: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "file is required: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	filename := header.Filename
	ext := strings.ToLower(filepath.Ext(filename))

	if ext != ".csv" {
		http.Error(w, "invalid file type: only .csv allowed", http.StatusBadRequest)
		return
	}

	data, err := utils.ParseCSV(file)
	if err != nil {
		http.Error(w, "failed to parse csv: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.svc.Upload(data); err != nil {
		http.Error(w, "validation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successfully uploaded transactions",
	})
}

func (h *TransactionHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	balance := h.svc.GetBalance()

	response := struct {
		Balance int64 `json:"balance"`
	}{
		Balance: balance,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *TransactionHandler) GetIssues(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	issues := h.svc.GetIssues()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(issues)
}
