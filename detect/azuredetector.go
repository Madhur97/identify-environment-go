package detect

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// BareMetalDetector checks for indications of a non-virtualized environment
type AzureDetector struct{}

type IMDSResponse struct {
	Compute struct {
		SubscriptionId string `json:"subscriptionId"`
	} `json:"compute"`
}

func (AzureDetector) GetIdentifiers() ([]string, error) {
	var identifiers []string
	client := &http.Client{}
	// Create a request to the Azure Instance Metadata Service (IMDS)
	req, err := http.NewRequest("GET", "http://169.254.169.254/metadata/instance?api-version=2021-02-01", nil)
	if err != nil {
		return identifiers, fmt.Errorf("failed to create request: %v", err)
	}

	// IMDS requires the "Metadata: true" header to be set for security reasons
	req.Header.Set("Metadata", "true")

	// Perform the HTTP GET request
	resp, err := client.Do(req)
	if err != nil {
		return identifiers, fmt.Errorf("failed to perform request: %v", err)
	}
	defer resp.Body.Close()

	// Read and return the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return identifiers, fmt.Errorf("failed to read response body: %v", err)
	}
	var data IMDSResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return identifiers, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	identifiers = append(identifiers, data.Compute.SubscriptionId)
	return identifiers, nil
}
