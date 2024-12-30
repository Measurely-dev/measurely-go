package measurely

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CapturePayload struct {
	Value int `json:"value"`
}

type CaptureResult struct {
	Success bool
	Message string
}

var API_KEY string = ""

func Init(NEW_API_KEY string) {
	API_KEY = NEW_API_KEY
}

func Capture(metric_identifier string, payload CapturePayload) CaptureResult {
	if API_KEY == "" {
		return CaptureResult{
			Success: false,
			Message: "Missing API KEY, please call the init function",
		}
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return CaptureResult{
			Success: false,
			Message: "Failed to marshal the payload",
		}
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("https://api.measurely.dev/event/v1/%s", metric_identifier), bytes.NewBuffer(body))
	if err != nil {
		return CaptureResult{
			Success: false,
			Message: "Failed to create http post request",
		}
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", API_KEY))
	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return CaptureResult{
			Success: false,
			Message: "Failed to send request",
		}
	}

	body, err = io.ReadAll(response.Body)
	if err != nil {
		return CaptureResult{
			Success: false,
			Message: "Failed to read response body",
		}
	}

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
