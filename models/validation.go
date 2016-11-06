package models

import (
	"encoding/json"
	"net/http"
)

type InputValidation interface {
	Validate(r *http.Request) error
}

func DecodeAndValidate(r *http.Request, v InputValidation) error {

	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}

	defer r.Body.Close()
	return v.Validate(r)
}
