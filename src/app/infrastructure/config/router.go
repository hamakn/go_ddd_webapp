package config

import (
	"github.com/gorilla/mux"
	"github.com/hamakn/go_ddd_webapp/src/app/interface/handler"
	"github.com/hamakn/go_ddd_webapp/src/app/interface/middleware"
	"github.com/urfave/negroni"
)

// NewRouter returns Negroni router to handle http request
func NewRouter() *negroni.Negroni {
	router := mux.NewRouter()

	// User
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/", handler.GetUsers()).Methods("GET")
	userRouter.HandleFunc("/{key}", handler.GetUser()).Methods("GET")

	n := negroni.New(
		middleware.NewContextSetter(),
	)
	n.UseHandler(router)

	return n
}
