package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestRegisterInvitation(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/invitations/{code}", RegisterInvitation).Methods("POST")

	body := []byte(`{"email": "test@example.com"}`)
	req, err := http.NewRequest("POST", "/api/v1/invitations/twitter-reg1", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")
}
