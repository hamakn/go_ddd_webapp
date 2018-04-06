package config

import (
	"github.com/gorilla/mux"
	"github.com/hamakn/go_ddd_webapp/src/app/interfaces/handler"
	"github.com/hamakn/go_ddd_webapp/src/app/interfaces/middleware"
	"github.com/urfave/negroni"
)

// NewRouter returns Negroni router to handle http request
func NewRouter() *negroni.Negroni {
	router := mux.NewRouter()

	// User
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/", handler.GetUsers()).Methods("GET")
	userRouter.HandleFunc("/{id}", handler.GetUser()).Methods("GET")
	userRouter.HandleFunc("/", handler.CreateUser()).Methods("POST")

	n := negroni.New(
		middleware.NewContextSetter(),
	)
	n.UseHandler(router)

	return n
}
