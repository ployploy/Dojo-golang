package main

import (
	"fmt"
	"net/http"
)

type HomePageHandler struct{}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello,ploy")
}
func main() {
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello,world")
		})
	barHandle := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello,bar")
	}
	http.Handle("/home", &HomePageHandler{})
	http.HandleFunc("/bar", barHandle)
	http.ListenAndServe(":3000", nil)
}
