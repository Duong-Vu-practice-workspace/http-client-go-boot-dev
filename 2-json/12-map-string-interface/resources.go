package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

func getResources(url string) ([]map[string]any, error) {
	var resources []map[string]any

	res, err := http.Get(url)
	if err != nil {
		return resources, err
	}

	defer res.Body.Close()

	// ?
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resources, err
	}
	var data []map[string]interface{}
	err = json.Unmarshal(body, &data)
	return data, nil

}

func logResources(resources []map[string]any) {
	var formattedStrings []string

	// ?
	for _, resource := range resources {
		for key := range resource {
			value, ok := resource[key]
			if ok {
				s := fmt.Sprintf("Key: %s - Value: %v", key, value)
				formattedStrings = append(formattedStrings, s)
			}
		}
	}
	sort.Strings(formattedStrings)

	for _, str := range formattedStrings {
		fmt.Println(str)
	}
}
