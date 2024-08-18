package common

import (
	"golang-template-api-service/app/internal/dto/common"
	"encoding/json"
	"testing"
)

func TestUploadResponse(t *testing.T) {
	// Create an instance of UploadResponse
	response := common.UploadResponse{
		Message:   "Upload successful",
		ObjectURL: "https://example.com/uploaded/file.jpg",
	}

	// Serialize the response to JSON
	jsonData, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Error marshaling UploadResponse to JSON: %v", err)
	}

	// Deserialize the JSON back into an UploadResponse struct
	var decodedResponse common.UploadResponse
	if err := json.Unmarshal(jsonData, &decodedResponse); err != nil {
		t.Fatalf("Error unmarshaling JSON to UploadResponse struct: %v", err)
	}

	// Check the fields of the decoded response
	if decodedResponse.Message != "Upload successful" {
		t.Error("Message should be 'Upload successful'")
	}

	if decodedResponse.ObjectURL != "https://example.com/uploaded/file.jpg" {
		t.Error("ObjectURL should be 'https://example.com/uploaded/file.jpg'")
	}
}
