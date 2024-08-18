package common

import (
	"golang-template-api-service/app/internal/dto/common"
	"encoding/json"
	"testing"
)

type APIError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func TestResponse(t *testing.T) {
	// Create an instance of Response
	response := common.Response{
		APIVersion: "1.0",
		Status:     "success",
		Data:       struct{ Key string }{Key: "value"},
	}

	// Serialize the response to JSON
	jsonData, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Error marshaling response to JSON: %v", err)
	}

	// Deserialize the JSON back into a Response struct
	var decodedResponse common.Response
	if err := json.Unmarshal(jsonData, &decodedResponse); err != nil {
		t.Fatalf("Error unmarshaling JSON to Response struct: %v", err)
	}

	// Check the fields of the decoded response
	if decodedResponse.APIVersion != "1.0" {
		t.Error("APIVersion should be '1.0'")
	}

	if decodedResponse.Status != "success" {
		t.Error("Status should be 'success'")
	}

	if decodedResponse.Data.(map[string]interface{})["Key"] != "value" {
		t.Error("Data field should contain 'Key' with value 'value'")
	}
}

func TestAPIError(t *testing.T) {
	// Create an instance of APIError
	apiError := APIError{
		Status:  "error",
		Message: "An error occurred",
	}

	// Serialize the API error to JSON
	jsonData, err := json.Marshal(apiError)
	if err != nil {
		t.Fatalf("Error marshaling API error to JSON: %v", err)
	}

	// Deserialize the JSON back into an APIError struct
	var decodedAPIError APIError
	if err := json.Unmarshal(jsonData, &decodedAPIError); err != nil {
		t.Fatalf("Error unmarshaling JSON to APIError struct: %v", err)
	}

	// Check the fields of the decoded API error
	if decodedAPIError.Status != "error" {
		t.Error("Status should be 'error'")
	}

	if decodedAPIError.Message != "An error occurred" {
		t.Error("Message should be 'An error occurred'")
	}
}
