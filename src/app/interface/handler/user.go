package handler

import (
	"net/http"

	"github.com/hamakn/go_ddd_webapp/src/app/interface/response"
)

var GetUsers = createAppHandler(func(w http.ResponseWriter, r *http.Request) (*response.Response, *appError) {
	return &response.Response{[]byte("this is response!!")}, nil
})
