package config

import (
	"github.com/gorilla/mux"
	"github.com/hamakn/go_ddd_webapp/src/app/interface/handler"
	"github.com/urfave/negroni"
)

// NewRouter returns Negroni router to handle http request
func NewRouter() *negroni.Negroni {
	router := mux.NewRouter()

	// User
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/", handler.GetUsers).Methods("GET")

	n := negroni.New()
	n.UseHandler(router)

	return n
}
