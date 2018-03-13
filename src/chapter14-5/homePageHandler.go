package test

import (
	"encoding/json"
	"net/http"
	"time"
)

type HomePageHandler struct{}
type User struct {
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	CreateAT  time.Time `json:"createat"`
}

func (h *HomePageHandler) Serve(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	json.NewDecoder(r.Body).Decode(user)
	user.CreateAT = time.Now()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	data, _ := json.Marshal(user)
	w.Write(data)
}
