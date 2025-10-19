package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http/httptest"

	"github.com/Nutan-Kum12/Ecom/types"
	"github.com/Nutan-Kum12/Ecom/utils"
)

func main() {
	// Test JSON parsing
	testJSON := `{
		"firstName": "John",
		"lastName": "Doe", 
		"email": "john@example.com",
		"password": "password123"
	}`

	fmt.Println("Testing JSON parsing...")
	fmt.Println("JSON to parse:", testJSON)

	// Create a test HTTP request
	req := httptest.NewRequest("POST", "/test", bytes.NewBufferString(testJSON))
	req.Header.Set("Content-Type", "application/json")

	// Test our ParseJSON function
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(req, &payload); err != nil {
		log.Printf("ParseJSON error: %v", err)
	} else {
		fmt.Printf("Parsed successfully: %+v\n", payload)
	}

	// Test standard json.Unmarshal
	var payload2 types.RegisterUserPayload
	if err := json.Unmarshal([]byte(testJSON), &payload2); err != nil {
		log.Printf("json.Unmarshal error: %v", err)
	} else {
		fmt.Printf("json.Unmarshal result: %+v\n", payload2)
	}

	// Test validation
	if err := utils.Validate.Struct(payload); err != nil {
		fmt.Printf("Validation error: %v\n", err)
	} else {
		fmt.Println("Validation passed!")
	}
}
