package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

type Response struct {
	Message string `json:"message"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	res := &Response{Message: "hello"}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	spew.Dump(r.Header)
	json.NewEncoder(w).Encode(res)
}
func main() {
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
