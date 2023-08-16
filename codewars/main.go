package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	clientID := "UWJidWxpeWJHdEJvZHNXd0ZvMFp4RG9ocjhzYTozdng2cG9wUjV1MDdWckNBNFJtYWVNRDNmX0Vh"
	accessTokenURL := "https://gateway.api.cloud.wso2.com/token"

	payload := []byte("grant_type=client_credentials")
	req, err := http.NewRequest("POST", accessTokenURL, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	authHeader := fmt.Sprintf("Basic %s", clientID)
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response:", string(body))
}
