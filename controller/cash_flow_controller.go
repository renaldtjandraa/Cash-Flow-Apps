package controller

import (
	config "Cash-Flow-Apps/config"
	model "Cash-Flow-Apps/model"
	"context"
	"encoding/json"
	"net/http"
)

func GetCashFlow(w http.ResponseWriter, r *http.Request) {
	// Koneksi ke database
	db, err := config.Connect()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close(context.Background())

	var response model.CashFlowResponse
	var cashFlows []model.CashFlow

	query := `SELECT * FROM cash_flow`
	rows, err := db.Query(context.Background(), query)
	if err != nil {
		response.Status = 400
		response.Message = "Error executing query"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var cashFlow model.CashFlow
		if err := rows.Scan(&cashFlow.ID, &cashFlow.Date, &cashFlow.Name, &cashFlow.Amount,
			&cashFlow.ExpenseType, &cashFlow.Description, &cashFlow.Status, &cashFlow.Comments); err != nil {
			response.Status = 500
			response.Message = "Error scanning data"
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}
		cashFlows = append(cashFlows, cashFlow)
	}

	if len(cashFlows) > 0 {
		response.Status = 200
		response.Message = "Success retrieving data"
		response.Data = cashFlows
	} else {
		response.Status = 404
		response.Message = "No data found"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
