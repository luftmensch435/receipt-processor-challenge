package main

import (
	"fmt"
	"net/http"
)

func main() {
	// add handlers
	http.HandleFunc("/receipts/process", HandleProcessReceipts)

	// start server
	fmt.Println("Server starting on port 8080....")
	fmt.Println("http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed: %v\n", err)
	}
}
