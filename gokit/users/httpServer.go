package users

import (
	"context"
	"encoding/json"
	goKitHttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHttpServer(ctx context.Context, endpoints *Endpoints) http.Handler {

	router := mux.NewRouter()
	router.Use(HeaderMiddleWare)
	router.Methods("POST").Path("/user").Handler(
		goKitHttp.NewServer(
			endpoints.CreateUser,
			decodeUserRequest,
			encodeUserResponse,
		))

	router.Methods("GET").Path("/user/{email}").Handler(
		goKitHttp.NewServer(
			endpoints.GetUser,
			decodeGetUserRequest,
			encodeGetUserRequest,
		))

	return router
}

func encodeGetUserRequest(ctx context.Context, writer http.ResponseWriter, i interface{}) error {
	return json.NewEncoder(writer).Encode(i)
}

func decodeGetUserRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	return CreateGetUserRequest{
		Email: vars["email"],
	}, nil
}

func encodeUserResponse(ctx context.Context, writer http.ResponseWriter, i interface{}) error {
	return json.NewEncoder(writer).Encode(i)
}

func decodeUserRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func HeaderMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(writer, request)
	})
}
