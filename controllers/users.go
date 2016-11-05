package controllers

import (
	"encoding/json"
	"github.com/anthify/api-slideshow-go/models"
	"io"
	"io/ioutil"
	"net/http"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	newUser := models.User{}
	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	json.Unmarshal(body, &newUser)
	user, _ := models.Create(newUser)

	RenderJson(w, user)
}
