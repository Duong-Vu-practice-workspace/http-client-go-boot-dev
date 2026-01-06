package main

import (
	"net/http"
)

func deleteUser(baseURL, id, apiKey string) error {
	fullURL := baseURL + "/" + id

	// ?
	request, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode > 299 {
		return err
	}
	return nil
}
