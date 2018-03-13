package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	http.Handle("/", &HomePageHandler{})
	http.ListenAndServe(":3000", nil)
}

type User struct {
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	CreateAt  time.Time `json:"createat"`
}

type HomePageHandler struct{}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	json.NewDecoder(r.Body).Decode(user)
	user.CreateAt = time.Now()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	data, _ := json.Marshal(user)
	w.Write(data)
}
