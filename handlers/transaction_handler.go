package handlers

import (
	"crud-category/models"
	"crud-category/services"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		service: service,
	}
}

func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request models.CheckoutRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "invalid request body",
		})
		return
	}

	transaction, err := h.service.Checkout(request.Items)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "something went wrong",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}

func (h *TransactionHandler) TransactionReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	start_date := r.URL.Query().Get("start_date")
	end_date := r.URL.Query().Get("end_date")

	if start_date == "" || end_date == "" {
		log.Println("start_date or end_date is missing")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "start_date and end_date are required",
		})
		return
	}

	if _, err := time.Parse("2006-01-02", start_date); err != nil {
		log.Println("invalid start_date format")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "invalid start_date format",
		})
		return
	}

	if _, err := time.Parse("2006-01-02", end_date); err != nil {
		log.Println("invalid end_date format")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "invalid end_date format",
		})
		return
	}

	if start_date > end_date {
		log.Println("start_date is greater than end_date")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "start_date cannot be greater than end_date",
		})
		return
	}

	transactions, err := h.service.TransactionReport(start_date, end_date)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": "something went wrong",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}
