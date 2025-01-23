package main

// in-memory storage to store receipts
// NOTE: Assumes single-threaded environment, mutex lock needed for production system.
var storage = make(map[string]Receipt)

// add receipt to storage
func SaveReceipt(id string, receipt Receipt) {
	storage[id] = receipt
}

// get receipt by id
func GetReceipt(id string) (Receipt, bool) {
	receipt, exists := storage[id]
	return receipt, exists
}
