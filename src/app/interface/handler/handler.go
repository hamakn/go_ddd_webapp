package handler

import (
	"net/http"

	"github.com/hamakn/go_ddd_webapp/src/app/interface/response"
)

type appError struct {
	Error   error
	Message string
	Code    int
}

func createAppHandler(f func(http.ResponseWriter, *http.Request) (*response.Response, *appError)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res, ae := f(w, r)

		w.Header().Set("Content-Type", "application/json")

		if ae != nil {
			// TODO
			return
		}

		w.Write(res.Body)
	}
}
