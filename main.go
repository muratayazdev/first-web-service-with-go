package main

import (
	"net/http"

	"muratayaz.com/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
