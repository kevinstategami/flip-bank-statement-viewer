package handler

import (
	"bytes"
	"encoding/json"
	"flip-bank-statement-viewer/internal/model"
	"flip-bank-statement-viewer/internal/repository"
	"flip-bank-statement-viewer/internal/service"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestHandler() *TransactionHandler {
	data := []model.Transaction{}
	repo := repository.NewTransactionRepository(&data)
	svc := service.NewTransactionService(repo)
	return NewTransactionHandler(svc)
}

func TestGetBalanceHandler(t *testing.T) {
	h := newTestHandler()

	req := httptest.NewRequest(http.MethodGet, "/balance", nil)
	w := httptest.NewRecorder()

	h.GetBalance(w, req)

	if w.Code != 200 {
		t.Fatalf("expected status 200 got %d", w.Code)
	}

	var resp map[string]int64
	json.Unmarshal(w.Body.Bytes(), &resp)

	if resp["balance"] != 0 {
		t.Fatalf("expected balance 0 got %d", resp["balance"])
	}
}

func TestUpload_InvalidExtension(t *testing.T) {
	h := newTestHandler()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	fileWriter, _ := writer.CreateFormFile("file", "test.txt")
	fileWriter.Write([]byte("dummy"))

	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()
	h.Upload(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 invalid file type, got %d", w.Code)
	}
}
