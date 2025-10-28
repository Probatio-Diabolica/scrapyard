package main

import (
	"net/http"
)

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	responseWithJson(w, 200, struct{}{})
}
