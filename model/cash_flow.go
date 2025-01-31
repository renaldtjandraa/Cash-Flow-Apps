package model

import "time"

type CashFlow struct {
	ID          int       `json:"id"`
	Date        time.Time `json:"date"`
	Name        string    `json:"name"`
	Amount      float64   `json:"amount"`
	ExpenseType string    `json:"expense_type"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Comments    string    `json:"comments"`
}

type CashFlowResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []CashFlow `json:"data,omitempty"`
}
