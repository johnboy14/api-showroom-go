package controllers

import (
	"fmt"
	"github.com/anthify/api-slideshow-go/models"
	"net/http"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	err := models.DecodeAndValidate(r, user)

	if err != nil {
		// send a bad request back to the caller
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		return
	}

	persistedUser, _ := models.Create(user)

	RenderJson(w, persistedUser)
}
