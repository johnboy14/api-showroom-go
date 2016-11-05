package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func HandlerFunc(w http.ResponseWriter, req *http.Request) {
	router := NewRouter()
	router.ServeHTTP(w, req)
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", UserCreate).Methods("PUT")
	return router
}

func RenderJson(rw http.ResponseWriter, object interface{}) {
	data, _ := json.MarshalIndent(object, "", "  ")

	data = append(data, '\n')

	rw.Header().Set("Content-Type", "application/json")

	rw.Write(data)
}
