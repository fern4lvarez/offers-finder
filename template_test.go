package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRenderMapTemplate(t *testing.T) {
	response := httptest.NewRecorder()

	RenderMapTemplate(response)

	if response.Code != 200 {
		t.Errorf(errorMessage, "RenderMapTemplate", 200, response.Code)
	}

	if !strings.HasPrefix(response.Body.String(), "<!DOCTYPE html>") {
		t.Errorf(errorMessage, "RenderMapTemplate",
			"<!DOCTYPE html>",
			"response.Body.String()")
	}
}
