package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/hamakn/go_ddd_webapp/src/app/interfaces/response"
	"google.golang.org/appengine/log"
	validator "gopkg.in/go-playground/validator.v9"
)

type appError struct {
	Error   error
	Message string
	Code    int
}

func createAppHandler(f func(context.Context, http.ResponseWriter, *http.Request) (*response.Response, *appError)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		res, apperr := f(ctx, w, r)

		w.Header().Set("Content-Type", "application/json")

		if apperr != nil {
			log.Errorf(ctx, "%#v", apperr.Error)

			if err := WriteErrorResponse(w, apperr.Code, apperr.Message); err != nil {
				log.Errorf(ctx, "%#v", errors.Wrap(apperr.Error, err.Error()))
			}

			return
		}

		if isEmptyResponse(res) {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		_, err := w.Write(res.Body)
		if err != nil {
			log.Errorf(ctx, "%#v", err)
		}
	}
}

// WriteErrorResponse is writing error response
func WriteErrorResponse(w http.ResponseWriter, code int, message string) error {
	w.WriteHeader(code)

	res, err := response.NewErrorResponse(message)
	if err != nil {
		return err
	}

	_, err = w.Write(res.Body)
	return err
}

func isEmptyResponse(res *response.Response) bool {
	if len(res.Body) == 0 {
		return true
	}
	return false
}

func parseRequest(r *http.Request, request interface{}) error {
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		return err
	}

	v := validator.New()
	return v.Struct(request)
}
