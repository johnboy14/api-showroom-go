package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Response struct {
	Errors string      `json:"errors"`
	Data   interface{} `json:"data"`
}

func BadRequestResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	response := Response{Errors: err.Error(), Data: nil}
	RenderJson(w, response)
}

func CreatedResponse(w http.ResponseWriter, object interface{}) {
	w.WriteHeader(http.StatusCreated)
	response := Response{Errors: "", Data: object}
	RenderJson(w, response)
}

func RenderJson(w http.ResponseWriter, object interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(object)
}

func DecodeJsonTo(bytes *bytes.Buffer, object interface{}) {
	json.NewDecoder(bytes).Decode(object)
}
