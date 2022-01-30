package main

import "fmt"

type ECODE int

type ErrorResponse struct {
	ErrorCode     ECODE  `json:"code"`
	ErrorMessage  string `json:"errorMessage"`
	ErrorDetails  string `json:"errorDetails"`
	TransactionId string `json:"transactionId"`
}

const (
	// Error codes
	//NOTE: These are not the same as the HTTP status codes.

	//NO ERROR
	ECODE_OK ECODE = iota

	//INVALID INPUT ERRORS
	ECODE_ENTITY_NOT_FOUND ECODE = iota + 1000
	ECODE_VALIDATION_ERROR
	ECODE_CONFLICT_ERROR

	//SYSTEM ERRORS
	ECODE_DATABASE_ERROR ECODE = iota + 2000

	//UNKNOWN ERRORS
	ECODE_UNKNOWN_ERROR ECODE = iota + 9000
)

func main() {
	var test_error = ErrorResponse{
		ErrorCode:     ECODE_OK,
		ErrorMessage:  "No Error",
		ErrorDetails:  "",
		TransactionId: "",
	}

	fmt.Printf("Code: %[1]q(%[1]d), Details %s\n", test_error.ErrorCode, test_error.ErrorMessage)
}
