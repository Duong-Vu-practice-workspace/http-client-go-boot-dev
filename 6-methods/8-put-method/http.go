package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id

	// ?
	body, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}
	request, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(body))
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
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return User{}, err
	}
	var res User
	err = json.Unmarshal(responseData, &res)
	if err != nil {
		return User{}, err
	}
	return res, nil
}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id

	// ?
	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return User{}, err
	}
	request.Header.Set("X-API-Key", apiKey)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return User{}, err
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return User{}, err
	}
	var user User
	err = json.Unmarshal(data, &user)
	if err != nil {
		return User{}, err
	}
	return user, nil

}
