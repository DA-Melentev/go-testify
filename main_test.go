package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenRequestIsCorrect(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code to be 200")
	assert.NotEmpty(t, responseRecorder.Body.String(), "Expected non-empty response body")
}

func TestMainHandlerWhenCityIsUnsupported(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=unknown", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Expected status code to be 400")
	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "Expected specific error message for unsupported city")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expectedResponse := strings.Join(cafeList["moscow"], ",")
	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code to be 200")
	assert.Equal(t, expectedResponse, responseRecorder.Body.String(), "Expected all cafes to be returned")
}
