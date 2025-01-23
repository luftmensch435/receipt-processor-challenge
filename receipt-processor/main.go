package main

import (
	"fmt"
	"net/http"
)

func main() {
	// start server
	fmt.Println("Server starting on port 8080....")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed: %v\n", err)
	}
}

