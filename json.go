package main

import (
	"encoding/json"
)

type JsonResponse map[string]interface{}

func (jsonResponse JsonResponse) String() string {
	b, err := json.Marshal(jsonResponse)
	if err != nil {
		return ""
	}

	return string(b)
}
