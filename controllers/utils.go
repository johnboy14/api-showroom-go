package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func UnmarshallJsonBody(r *http.Request, o interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return error(err)
	}

	jsonerr := json.Unmarshal(body, &o)
	if jsonerr != nil {
		return error(err)
	}

	return nil
}

func RenderJson(rw http.ResponseWriter, object interface{}) {
	data, _ := json.MarshalIndent(object, "", "  ")

	data = append(data, '\n')

	rw.Header().Set("Content-Type", "application/json")

	rw.Write(data)
}
