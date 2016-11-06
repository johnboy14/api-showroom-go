package controllers

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSpec(t *testing.T) {

	Convey("User Controller Tests", t, func() {
		Convey("Given a valid User, Then return new user and 201 Status", func() {
			newUser := `{"first_name": "John", "last_name": "Ervine", "email": "johnervine@hotmail.com", "password": "password1"}`
			reader := strings.NewReader(string(newUser))

			req, _ := http.NewRequest("PUT", "/users", reader)
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(UserCreate)
			handler.ServeHTTP(rr, req)

			So(rr.Code, ShouldEqual, 201)
			So(rr.Header().Get("Content-Type"), ShouldEqual, "application/json")
		})
		Convey("Given a user, When email is invalid, Then return 400 with suitable Error message", func() {
			newUser := `{"first_name": "John", "last_name": "Ervine", "email": "johnervinehotmail.com", "password": "password1"}`
			reader := strings.NewReader(string(newUser))

			req, _ := http.NewRequest("PUT", "/users", reader)
			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(UserCreate)
			handler.ServeHTTP(rr, req)

			So(rr.Code, ShouldEqual, 400)
		})
	})
}
