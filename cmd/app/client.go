package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const (
	maxTokens = 150
)

func GenerateGPTText(query string) (string, error) {
	req := RequestGPT{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{
				Role:    "user",
				Content: query,
			},
		},
		MaxTokens: maxTokens,
	}

	reqJson, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	payload := bytes.NewBuffer(reqJson)

	request, err := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/completions", payload)
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer API_KEY")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var resp ResponseGPT
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return "", err
	}

	content := resp.Choices[0].Message.Content

	return content, nil
}
