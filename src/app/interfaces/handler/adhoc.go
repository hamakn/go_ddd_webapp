package handler

import (
	"net/http"

	"github.com/pkg/errors"
	"google.golang.org/appengine/log"
)

// GetError1 is test func for Stackdriver Logging without stacktrace
func GetError1() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		w.Header().Set("Content-Type", "application/json")

		err := errors.New("error1")
		err = errors.Wrap(err, "error2")

		// not show stacktrace
		log.Errorf(ctx, "%s", err)

		w.WriteHeader(http.StatusNoContent)
	}
}

// GetError2 is test func for Stackdriver Logging with stacktrace
func GetError2() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		w.Header().Set("Content-Type", "application/json")

		err := errors.New("error1")
		err = errors.Wrap(err, "error2")

		// show stacktrace
		log.Errorf(ctx, "%+v", err)

		w.WriteHeader(http.StatusNoContent)
	}
}

// GetPanic1 is test func for Stackdriver Logging with panic
func GetPanic1() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		panic("panic1")
	}
}
