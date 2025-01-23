package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"

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
	if receipt.Retailer == "" ||
		receipt.PurchaseDate == "" ||
		receipt.PurchaseTime == "" ||
		receipt.Total == "" ||
		len(receipt.Items) == 0 {
		return false
	}

	for _, item := range receipt.Items {
		if item.ShortDescription == "" || item.Price == "" {
			return false
		}
	}

	return true
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
		http.Error(w, "The receipt is invalid.", http.StatusBadRequest)
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
	points := 0

	// One point for every alphanumeric character in the retailer name.
	for _, r := range receipt.Retailer {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			points += 1
		}
	}

	// 50 points if the total is a round dollar amount with no cents.
	// NOTE: Assumes `total` is a valid value that can be parsed into float64.
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if math.Mod(total, 1) == 0 {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25.
	// NOTE: Assumes `total` is a valid value that can be parsed into float64.
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// 5 points for every two items on the receipt.
	points += 5 * (len(receipt.Items) / 2)

	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	// NOTE: Assumes `price` is a valid value that can be parsed into float64.
	for _, r := range receipt.Items {
		descriptionLength := len(strings.TrimSpace(r.ShortDescription))
		if descriptionLength%3 == 0 {
			itemPrice, _ := strconv.ParseFloat(r.Price, 64)
			points += int(math.Ceil(itemPrice * 0.2))
		}
	}

	// 6 points if the day in the purchase date is odd.
	// NOTE: Assumes `purchaseDate` is a valid value that can be parsed into this time layout "2022-01-01".
	purchaseDate, _ := time.Parse("2022-01-01", receipt.PurchaseDate)
	if purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	// NOTE: Assumes `purchaseTime` is a valid value that can be parsed into this time layout "13:01".
	purchaseTime, _ := time.Parse("13:01", receipt.PurchaseTime)
	if purchaseTime.Hour() == 14 || purchaseTime.Hour() == 15 {
		points += 10
	}

	return strconv.Itoa(points)
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
		http.Error(w, "No receipt found for that ID.", http.StatusNotFound)
		return
	}

	// Calculate points
	points := calculatePoints(receipt)

	// Return points
	responseData := map[string]string{"points": points}
	encoder := json.NewEncoder(w)
	encoder.Encode(responseData)
}
