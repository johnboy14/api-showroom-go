package controllers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Errors string      `json:"errors"`
	Data   interface{} `json:"data"`
}

func BadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	response := Response{Errors: err.Error()}
	RenderJson(w, response)
}

func Created(w http.ResponseWriter, object interface{}) {
	w.WriteHeader(http.StatusCreated)
	response := Response{Errors: "", Data: object}
	RenderJson(w, response)
}

func RenderJson(w http.ResponseWriter, object interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(object)
}
