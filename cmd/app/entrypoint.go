package main

import (
	"encoding/base64"
	"errors"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
)

func Process(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	result, err := base64Decoder(request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	text, err := GenerateGPTText(result)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       text,
	}, nil
}

func base64Decoder(request string) (string, error) {
	dataBytes, err := base64.StdEncoding.DecodeString(request)
	if err != nil {
		return "", err
	}

	data, err := url.ParseQuery(string(dataBytes))
	if err != nil {
		return "", err
	}

	if data.Has("Body") {
		return data.Get("Body"), nil
	}

	return "", errors.New("body not found")
}
