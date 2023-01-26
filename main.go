package main

import (
	"bytes"
	"encoding/json"
	s "example/configs"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Replace YOUR_API_KEY with your actual API key
	apiKey := s.API_KEY
	url := "https://api.openai.com/v1/engines/davinci-codex/completions"

	// create the JSON payload
	payload := map[string]interface{}{
		"prompt":      "What is the capital city of France?",
		"temperature": 0.5,
	}
	jsonValue, _ := json.Marshal(payload)

	// create the request
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey)

	// make the request
	client := &http.Client{}
	resp, _ := client.Do(req)

	// read the response
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
