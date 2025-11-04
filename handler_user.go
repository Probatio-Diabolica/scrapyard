package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"scrapyard/internal/auth"
	"scrapyard/internal/database"
	"time"

	"github.com/google/uuid"
)

func (api *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	par := params{}

	err := decoder.Decode(&par)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("error parsing the json: %s", err))
		return
	}

	usr, err := api.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      par.Name,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("error creating user: %s", err))
		return
	}

	responseWithJson(w, 201, databaseUserToUser(usr))
}

func (api *apiConfig) handleGetUserByAPIKey(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		responseWithError(w, 403, fmt.Sprintf("auth error: %v", err))
		return
	}

	user, err := api.DB.GetUserByAPIKey(r.Context(), apiKey)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("user unfetchable: %v", err))
		return
	}

	responseWithJson(w, 200, databaseUserToUser(user))
}
