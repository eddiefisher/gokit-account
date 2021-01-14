package account

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/auth/jwt"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer realized server over HTTP, registered routers
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(
		commonMiddleware,
		jwtMiddleware,
	)
	options := []httptransport.ServerOption{
		httptransport.ServerBefore(jwt.HTTPToContext()),
	}

	r.Methods(http.MethodPost).Path("/user").Handler(httptransport.NewServer(
		endpoints.CreateUser,
		decodeUserRequest,
		encodeResponse,
		options...,
	))

	r.Methods(http.MethodGet).Path("/user/{uuid}").Handler(httptransport.NewServer(
		endpoints.GetUser,
		decodeEmailRequest,
		encodeResponse,
		options...,
	))

	return r
}
