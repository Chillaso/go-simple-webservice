package main

import (
	"net/http"

	"github.com/chillaso/go-webservice/controller"
)

func main() {
	controller.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
