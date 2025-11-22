package utils

import (
	"strings"
	"testing"
)

func TestParseCSV_Valid(t *testing.T) {
	input := `
1624507883,JOHN DOE,DEBIT,250000,SUCCESS,restaurant
1624608050,E-COMMERCE A,DEBIT,150000,FAILED,clothes
`
	reader := strings.NewReader(input)

	res, err := ParseCSV(reader)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(res) != 2 {
		t.Fatalf("expected 2 rows, got %d", len(res))
	}

	if res[0].Amount != 250000 {
		t.Fatalf("expected amount 250000 got %d", res[0].Amount)
	}
}

func TestParseCSV_InvalidColumnCount(t *testing.T) {
	input := `1624507883,JOHN DOE,DEBIT,250000`

	reader := strings.NewReader(input)
	_, err := ParseCSV(reader)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestParseCSV_InvalidType(t *testing.T) {
	input := `1624507883,JOHN DOE,INVALID,250000,SUCCESS,restaurant`
	_, err := ParseCSV(strings.NewReader(input))
	if err == nil {
		t.Fatalf("expected error for invalid type")
	}
}
