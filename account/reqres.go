package account

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateUserRequest ...
type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUserResponse ...
type CreateUserResponse struct {
	Ok string `json:"ok"`
}

// GetUserRequest ...
type GetUserRequest struct {
	UUID string `json:"uuid"`
}

// GetUserResponse ...
type GetUserResponse struct {
	Email string `json:"email"`
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeEmailRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetUserRequest
	vars := mux.Vars(r)
	req = GetUserRequest{
		UUID: vars["uuid"],
	}
	return req, nil
}
