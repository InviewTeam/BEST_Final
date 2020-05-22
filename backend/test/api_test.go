package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/InviewTeam/BEST_Final/api"
	"github.com/InviewTeam/BEST_Final/database"
)

var token string

func TestRegister(t *testing.T) {
	creds := database.Credentials{
		Username: "someUser",
		Password: "secretpasswd",
	}
	reqBody, _ := json.Marshal(creds)

	req, err := http.NewRequest("POST", "/register", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.Register)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v; %s",
			status, http.StatusOK, string(reqBody))
	}
}

/*
func TestLogin(t *testing.T) {
	creds := database.Credentials{
		Username: "someUser",
		Password: "secretpasswd",
	}
	reqBody, _ := json.Marshal(creds)

	req, err := http.NewRequest("POST", "/login", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.Register)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var tk database.Token
	bodyBytes, _ := ioutil.ReadAll(rr.Body)
	err = json.Unmarshal(bodyBytes, &tk)
	if err != nil {
		t.Errorf("failed to parse body: %s", rr.Body.String())
	} else {
		token = tk.Token
	}*/
