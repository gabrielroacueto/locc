package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const MODEL = "llama3.2"

func Generate(prompt string) (string, error) {
	resp, err := ollamaGenerateRequest(prompt, false)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	var response struct {
		Response string `json:"response"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("error decoding response: %w", err)
	}

	return response.Response, nil
}

func GenerateStream(prompt string, callback func(string)) error {
	resp, err := ollamaGenerateRequest(prompt, true)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read response line by line
	decoder := json.NewDecoder(resp.Body)
	for {
		// Response structure for each chunk
		var response struct {
			Response string `json:"response"`
			Done     bool   `json:"done"`
		}

		if err := decoder.Decode(&response); err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error decoding response: %w", err)
		}

		// Send chunk to callback
		callback(response.Response)

		// Break if we're done
		if response.Done {
			break
		}
	}

	return nil
}

func ollamaGenerateRequest(prompt string, stream bool) (*http.Response, error) {
	reqBody := struct {
		Model  string `json:"model"`
		Prompt string `json:"prompt"`
		Stream bool   `json:"stream"`
	}{
		Model:  MODEL,
		Prompt: prompt,
		Stream: stream,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	// Create request
	req, err := http.NewRequest("POST", "http://localhost:11434/api/generate", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	return client.Do(req)
}
