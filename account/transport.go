package account

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer realized server over HTTP, registered routers
func NewHTTPServer(ctx context.Context, endpoints Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	r.Use(
		commonMiddleware,
		jwtMiddleware,
	)
	options := []httptransport.ServerOption{
		httptransport.ServerBefore(jwt.HTTPToContext()),
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
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

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
}
