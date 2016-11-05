package controllers

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", UserCreate).Methods("PUT")
	return router
}
