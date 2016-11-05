package main

import (
	"fmt"
	"github.com/anthify/api-slideshow-go/controllers"
	"github.com/codegangsta/negroni"
	"os"
)

var port = "8000"

func startWeb() {
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	n := negroni.Classic()
	n.UseHandler(controllers.NewRouter())
	n.Run(fmt.Sprintf(":%s", port))

}
