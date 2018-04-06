package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hamakn/go_ddd_webapp/src/app/application"
	"github.com/hamakn/go_ddd_webapp/src/app/application/request"
	"github.com/hamakn/go_ddd_webapp/src/app/interfaces/response"
	"github.com/pkg/errors"
)

var (
	// ErrGetUsers is error on GetUsers
	ErrGetUsers = errors.New("app-interface-handler-get-users: GetUsers failed")
	// ErrGetUser is error on GetUser
	ErrGetUser = errors.New("app-interface-handler-get-user: GetUser failed")
	// ErrCreateUser is error on CreateUser
	ErrCreateUser = errors.New("app-interface-handler-create-user: CreateUser failed")
)

// GetUsers is handler to handle getting users request
func GetUsers() func(http.ResponseWriter, *http.Request) {
	return createAppHandler(func(w http.ResponseWriter, r *http.Request) (*response.Response, *appError) {
		users, err := application.GetUsers(r.Context())
		if err != nil {
			return nil, &appError{errors.Wrap(err, ErrGetUsers.Error()), "internal server error", http.StatusInternalServerError}
		}

		res, err := response.UsersResponse(users)
		if err != nil {
			return nil, &appError{errors.Wrap(err, ErrGetUsers.Error()), "internal server error", http.StatusInternalServerError}
		}

		return res, nil
	})
}

// GetUser is handler to handle getting user request
func GetUser() func(http.ResponseWriter, *http.Request) {
	return createAppHandler(func(w http.ResponseWriter, r *http.Request) (*response.Response, *appError) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			return nil, &appError{err, "params id was wrong", http.StatusBadRequest}
		}

		u, err := application.GetUserByID(r.Context(), id)
		if err != nil {
			if err.Error() == "datastore: no such entity" {
				return nil, &appError{errors.Wrap(err, ErrGetUser.Error()), "Not Found", http.StatusNotFound}
			}
			return nil, &appError{errors.Wrap(err, ErrGetUser.Error()), "internal server error", http.StatusInternalServerError}
		}

		res, err := response.UserResponse(u)
		if err != nil {
			return nil, &appError{errors.Wrap(err, ErrGetUser.Error()), "internal server error", http.StatusInternalServerError}
		}

		return res, nil
	})
}

// CreateUser is handler to handle create user request
func CreateUser() func(http.ResponseWriter, *http.Request) {
	return createAppHandler(func(w http.ResponseWriter, r *http.Request) (*response.Response, *appError) {
		req := request.CreateUserRequest{}
		err := parseRequest(r, &req)
		if err != nil {
			return nil, &appError{errors.Wrap(err, ErrCreateUser.Error()), "Bad Request", http.StatusBadRequest}
		}

		u, err := application.CreateUser(r.Context(), req)
		if err != nil {
			return nil, &appError{errors.Wrap(err, ErrCreateUser.Error()), "internal server error", http.StatusInternalServerError}
		}

		res, err := response.UserResponse(u)
		if err != nil {
			return nil, &appError{errors.Wrap(err, ErrCreateUser.Error()), "internal server error", http.StatusInternalServerError}
		}

		return res, nil
	})
}
