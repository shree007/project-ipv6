package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRootEndpoint(t *testing.T) {
	expectedResponse := gin.H{
		"id":      "1",
		"message": "dlrow olleH",
	}
	router := SetupRouter()
	w := performRequest(router, "GET", "/?token=YXBwbGljYXRpb250d29fc2VjdXJlX3Rva2Vu")
	assert.Equal(t, http.StatusOK, w.Code)

	var actualResponse map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse["message"], actualResponse["message"])
	assert.Equal(t, expectedResponse["id"], actualResponse["id"])
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
