# Receipt Processor Submission
**Name: Yixin Tang (thomas.yixin@gmail.com)**

## Design Choices

### 1. Google's UUID Package
- Choice: [Google's Go package for UUIDs](https://github.com/google/uuid) for generating receipt IDs.
- Reason: While this is my first time writing anything in Go, I am already experienced with UUID usage in Python. Existing package offers safety, protocol compliance, scalability, and simplifies ID management by removing the need for complex custom code.
  
### 2. Testing Framework
- Choice: Python with `requests` library.
- Reason: I have extensive experience writing tests and making HTTP requests in Python. Since testing is optional for this project, leveraging a language I'm most comfortable with ensures efficiency and focus. 
  
## Project Structure
```plaintext
.
├── receipt-processor       # PROJECT SUBMISSION!!
│   ├── main.go             # Entry point for the Go server
│   ├── handler.go          # Handlers for API endpoints
│   └── storage.go          # In-memory data storage implementation
├── receipt-processor-test  # Contains Python test framework
│   ├── test_server.py      # Python script for testing 
│   └── test-receipts       # Test receipt JSON files
├── Instruction.md          # Old README.md
└── README.md               # Documentation
```

## Getting Started

### Prerequisites
- Go: Version 1.23 or higher.

### Steps to Run
```bash
cd receipt-process
go run .
```
The server will listen on http://localhost:8080.

### Testing

1. Ensure the Go server is running:  
   ```bash
   go run .
   ```

2. Run the test script:  
   ```bash
   cd receipt-process-test
   pip install requests
   python3 test_server.py
   ```

   **Note**: Due to limited computer resources, I am unable to run Docker on my system. As a result, a Dockerfile is not provided for testing. However, since Docker is not a requirement for this project, the testing can be performed directly using the provided Python script.
