package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// add handlers
	http.HandleFunc("/receipts/process", HandleProcessReceipts)
	http.HandleFunc("/receipts/", func(w http.ResponseWriter, r *http.Request) {
		// validate '/receipts/{id}/points'
		if strings.HasSuffix(r.URL.Path, "/points") {
			HandleGetPoints(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	// start server
	fmt.Println("Server starting on port 8080....")
	fmt.Println("http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed: %v\n", err)
	}
}
