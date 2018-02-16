package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hamakn/go_ddd_webapp/src/app/interface/response"
)

var (
	// ErrGetUsers is error on GetUsers
	ErrGetUsers = errors.New("app-interface-handler-get-users: GetUsers failed")
	// ErrGetUser is error on GetUser
	ErrGetUser = errors.New("app-interface-handler-get-user: GetUser failed")
)

// GetUsers is handler to handle getting users request
var GetUsers = createAppHandler(func(w http.ResponseWriter, r *http.Request) (*response.Response, *appError) {
	return &response.Response{[]byte("this is response!!")}, nil
})

// GetUser is handler to handle getting user request
var GetUser = createAppHandler(func(w http.ResponseWriter, r *http.Request) (*response.Response, *appError) {
	vars := mux.Vars(r)
	return nil, &appError{ErrGetUser, fmt.Sprintf("key: %v not found", vars["key"]), http.StatusNotFound}
})
