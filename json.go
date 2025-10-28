package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		//just some non-client side errors
		log.Println("Responding with 5xx error")
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	responseWithJson(w, code, errorResponse{
		Error: msg,
	})
}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("failed to marshal json response : ", payload)
		w.WriteHeader(500)
		return
	}

	//responding with json data
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
