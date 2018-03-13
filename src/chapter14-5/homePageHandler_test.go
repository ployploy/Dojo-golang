package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJsonHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/json", string.NewRecorder(`{"firstname" : "ploy", "lastname": "kiki", "email":"chonnikan@gmail.com"}`))

	home := HomePageHandler{}
	home.ServeHTTP(res, req)

	if res.Code != 200 {
		t.Fatalf("Expected status too === 200 , but got %d", res.Code)
	}

	user := new(User)
	json.NewDecoder(res.Body).Decode(user)

	if user.FirstName != "ploy" {
		t.Fatalf("Expected firstname too == ploy, but got %s", user.FirstName)
	}

}
