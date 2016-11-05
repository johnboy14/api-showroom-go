package controllers

import (
	"fmt"
	"github.com/anthify/api-slideshow-go/models"
	"net/http"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	newUser := models.User{}
	err := UnmarshallJsonBody(r, &newUser)

	if err != nil {
		fmt.Println(err)
	}

	user, _ := models.Create(newUser)

	RenderJson(w, user)
}
