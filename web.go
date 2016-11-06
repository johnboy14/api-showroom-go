package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/johnboy14/api-slideshow-go/controllers"
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
