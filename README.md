# Measurely-go

Measurely-go is a lightweight library for interacting with the Measurely API, enabling developers to manage and track custom metrics programmatically using Golang.

## Installation

To use the Measurely package in your Golang project, you need to import it and initialize it with your Measurely API key.

```bash
go get github.com/measurely-dev/measurely-go
```

## Usage

### 1. Initialize the Measurely package

Before you can send any metrics to Measurely, you need to initialize the package with your API key. The `Init` function accepts your API key as a string and sets it for use in subsequent API calls.

```go
package main

import (
	"github.com/measurely-dev/measurely-go"
	"fmt"
)

func main() {
	// Initialize the Measurely package with your API key
	measurely.Init("YOUR_API_KEY")
}
```

### 2. Send a Metric (Capture)

The `Capture` function is used to send metric data to Measurely. You need to pass the metric identifier (which is a unique name or ID for the metric) and the metric value you want to track.

#### Example of sending a metric

```go
package main

import (
	"github.com/measurely-dev/measurely-go"
	"fmt"
)

func main() {
	// Initialize the Measurely package with your API key
	measurely.Init("YOUR_API_KEY")

	// Create a metric payload
	payload := measurely.CapturePayload{
		Value: 42,
	}

	// Capture the metric and get the result
	result := measurely.Capture("example_metric", payload)

	// Handle the result
	if result.Success {
		fmt.Println("Metric captured successfully!")
	} else {
		fmt.Println("Error capturing metric:", result.Message)
	}
}
```

### 3. Error Handling

The `Capture` function returns a `CaptureResult` struct that contains two fields:

- `Success` (bool): Indicates if the API call was successful.
- `Message` (string): Contains the response message from the server, which could either be a success message or an error message.

### API Reference

#### `Init(NEW_API_KEY string)`

- **Description**: Initializes the Measurely package with your API key.
- **Parameters**:
  - `NEW_API_KEY`: The API key provided by Measurely.
- **Returns**: None.

#### `Capture(metric_identifier string, payload CapturePayload) CaptureResult`

- **Description**: Sends a metric value to Measurely for tracking.
- **Parameters**:
  - `metric_identifier`: The unique identifier for the metric you are capturing.
  - `payload`: A `CapturePayload` struct that contains the metric value to be recorded.
- **Returns**: A `CaptureResult` struct that contains the success status and response message.

### Types

#### `CapturePayload`

```go
type CapturePayload struct {
    Value int `json:"value"` // The metric value to be recorded.
}
```

- **Description**: This struct defines the data payload that is sent to the Measurely API when capturing a metric.
- **Fields**:
  - `Value` (int): The metric value that you want to track.

#### `CaptureResult`

```go
type CaptureResult struct {
    Success bool   // Indicates if the API call was successful.
    Message string // Contains the server's response or an error message.
}
```

- **Description**: This struct represents the result of the API call to capture a metric.
- **Fields**:
  - `Success` (bool): Indicates if the metric capture was successful.
  - `Message` (string): Contains the server's response or an error message.
