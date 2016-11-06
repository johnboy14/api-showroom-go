package controllers

import (
	"github.com/johnboy14/api-slideshow-go/models"
	"net/http"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	var user *models.User = &models.User{}

	err := models.DecodeAndValidate(r, user)
	if err != nil {
		BadRequestResponse(w, err)
		return
	}
	persistedUser, _ := models.Create(*user)
	CreatedResponse(w, persistedUser)
}
