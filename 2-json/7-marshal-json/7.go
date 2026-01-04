package __marshal_json

import (
	"encoding/json"
	"errors"
)

func marshalAll[T any](items []T) ([][]byte, error) {
	var res [][]byte

	for _, item := range items {
		data, err := json.Marshal(item)
		if err != nil {
			return nil, errors.New("cannot marshal items")
		}
		res = append(res, data)
	}

	return res, nil
}
