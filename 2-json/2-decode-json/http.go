package __decode_json

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	// ?
	issues := make([]Issue, 1)
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&issues); err != nil {
		return nil, fmt.Errorf("error decoding body")
	}
	return issues, nil
}
