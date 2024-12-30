package measurely

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// CapturePayload defines the structure of the data sent to the Measurely API.
type CapturePayload struct {
	Value int `json:"value"` // The metric value to be recorded.
}

// CaptureResult defines the structure of the result returned by the Measurely API.
type CaptureResult struct {
	Success bool   // Indicates if the API call was successful.
	Message string // Contains the server's response or an error message.
}

// API_KEY holds the API key for authenticating with the Measurely API.
var API_KEY string = ""

// Init initializes the Measurely library with the provided API key.
// @param NEW_API_KEY - The API key provided by Measurely.
func Init(NEW_API_KEY string) {
	API_KEY = NEW_API_KEY
}

// Capture sends a metric to the Measurely API.
// @param metric_identifier - The unique identifier or name of the metric to be tracked.
// @param payload - The data payload containing the metric value.
// @return CaptureResult - The result of the API call.
func Capture(metric_identifier string, payload CapturePayload) CaptureResult {
	// Ensure the API key is set before making the API call.
	if API_KEY == "" {
		return CaptureResult{
			Success: false,
			Message: "Missing API KEY, please call the init function",
		}
	}

	// Serialize the payload into JSON format.
	body, err := json.Marshal(payload)
	if err != nil {
		return CaptureResult{
			Success: false,
			Message: "Failed to marshal the payload",
		}
	}

	// Create the HTTP POST request with the metric data.
	request, err := http.NewRequest("POST", fmt.Sprintf("https://api.measurely.dev/event/v1/%s", metric_identifier), bytes.NewBuffer(body))
	if err != nil {
		return CaptureResult{
			Success: false,
			Message: "Failed to create http post request",
		}
	}

	// Set the headers for authorization and content type.
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", API_KEY))
	request.Header.Set("Content-Type", "application/json")

	// Send the HTTP request and capture the response.
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return CaptureResult{
			Success: false,
			Message: "Failed to send request",
		}
	}

	// Read the response body.
	body, err = io.ReadAll(response.Body)
	if err != nil {
		return CaptureResult{
			Success: false,
			Message: "Failed to read response body",
		}
	}

	// Convert the response body to a string and check for success.
	text := string(body)
	success := false
	if response.StatusCode == 200 {
		success = true
	}

	return CaptureResult{
		Success: success,
		Message: text,
	}
}

