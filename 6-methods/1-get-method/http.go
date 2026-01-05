package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func getUsers(url string) ([]User, error) {
	// ?
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(response.Body)
	var res []User
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}
