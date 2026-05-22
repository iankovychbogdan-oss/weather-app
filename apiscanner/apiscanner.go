package apiscanner

import (
	"fmt"
	"os"
	"strings"
)

func GetApiKey() string {
	data, err := os.ReadFile("apikey.txt")
	if err == nil {
		key := strings.TrimSpace(string(data))
		if key != "" {
			return key
		}
	}

	var apiKey string
	fmt.Print("Enter the API Key: ")
	fmt.Scan(&apiKey)

	os.WriteFile("apikey.txt", []byte(apiKey), 0600)

	return apiKey
}
