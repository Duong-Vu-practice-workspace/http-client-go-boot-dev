package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func createUser(url, apiKey string, data User) (User, error) {
	// ?
	requestBody, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return User{}, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return User{}, err
	}
	defer response.Body.Close()
	var user User
	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
