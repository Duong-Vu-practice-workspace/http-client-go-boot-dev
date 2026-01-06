package __error_handling

import (
	"fmt"
	"net/http"
)

func fetchData(url string) (int, error) {
	// ?
	response, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("network error: %v", err)
	}
	statusCode := response.StatusCode
	if statusCode != http.StatusOK {
		return response.StatusCode, fmt.Errorf("non-OK HTTP status: %s", response.Status)
	}
	return statusCode, nil
}
