package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

// receipt
type Receipt struct {
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Items        []Item
	Total        string
}

// receipt item
type Item struct {
	ShortDescription string
	Price            string
}

// validate receipt values, return True for valid
func validateReceipt(receipt Receipt) bool {
	return receipt.Retailer != "" &&
		receipt.PurchaseDate != "" &&
		receipt.PurchaseTime != "" &&
		receipt.Total != "" &&
		len(receipt.Items) > 0
}

// handle receipt submission
func HandleProcessReceipts(w http.ResponseWriter, r *http.Request) {
	// verify http method
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var receipt Receipt

	// decode, verify receipt
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&receipt)
	if err != nil || !validateReceipt(receipt) {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// generate ID, save
	id := uuid.New().String()
	SaveReceipt(id, receipt)

	// return ID
	responseData := map[string]string{"id": id}
	encoder := json.NewEncoder(w)
	encoder.Encode(responseData)
}

// calculate points according to rules
func calculatePoints(receipt Receipt) string {
	// points := 0
	// TODO: calculation
	return "I love money"
}

// retrieve points for a receipt
func HandleGetPoints(w http.ResponseWriter, r *http.Request) {
	// verify http method
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// extract ID from path '/receipts/{id}/points'
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	id := parts[1]

	// fetch receipt from memory
	receipt, exists := GetReceipt(id)
	if !exists {
		http.Error(w, "ID not exist", http.StatusNotFound)
		return
	}

	// Calculate points
	points := calculatePoints(receipt)

	// Return points
	responseData := map[string]string{"points": points}
	encoder := json.NewEncoder(w)
	encoder.Encode(responseData)
}
