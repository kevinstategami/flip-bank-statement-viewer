package utils

import (
	"encoding/csv"
	"errors"
	"flip-bank-statement-viewer/internal/model"
	"io"
	"strconv"
	"strings"
)

func ParseCSV(reader io.Reader) ([]model.Transaction, error) {
	csvReader := csv.NewReader(reader)
	csvReader.FieldsPerRecord = -1

	var transactions []model.Transaction
	rowIndex := 1

	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if len(row) == 0 || (len(row) == 1 && strings.TrimSpace(row[0]) == "") {
			rowIndex++
			continue
		}

		if len(row) != 6 {
			return nil, errors.New(
				"invalid CSV format at row " + strconv.Itoa(rowIndex) +
					": expected 6 columns, got " + strconv.Itoa(len(row)),
			)
		}

		for i := range row {
			row[i] = strings.TrimSpace(row[i])
		}

		ts, err := strconv.ParseInt(row[0], 10, 64)
		if err != nil {
			return nil, errors.New("invalid timestamp at row " + strconv.Itoa(rowIndex))
		}

		amount, err := strconv.ParseInt(row[3], 10, 64)
		if err != nil || amount < 0 {
			return nil, errors.New("invalid amount at row " + strconv.Itoa(rowIndex))
		}

		tType := model.TransactionType(row[2])
		if tType != model.Debit && tType != model.Credit {
			return nil, errors.New("invalid type at row " + strconv.Itoa(rowIndex) +
				": must be DEBIT or CREDIT")
		}

		tStatus := model.TransactionStatus(row[4])
		if tStatus != model.Success && tStatus != model.Failed && tStatus != model.Pending {
			return nil, errors.New("invalid status at row " + strconv.Itoa(rowIndex) +
				": must be SUCCESS, FAILED, or PENDING")
		}

		transaction := model.Transaction{
			Timestamp:   ts,
			Name:        row[1],
			Type:        tType,
			Amount:      amount,
			Status:      tStatus,
			Description: row[5],
		}

		transactions = append(transactions, transaction)
		rowIndex++
	}

	return transactions, nil
}
