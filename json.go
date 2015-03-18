package main

import (
	"encoding/json"
)

// JsonResponse provides a type for http json responses
type JsonResponse map[string]interface{}

// String formats the output of a JsonResponse
func (jsonResponse JsonResponse) String() string {
	b, err := json.Marshal(jsonResponse)
	if err != nil {
		return ""
	}

	return string(b)
}
