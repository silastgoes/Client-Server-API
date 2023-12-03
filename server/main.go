package main

import (
	"net/http"

	"github.com/silastgoes/client-server-api/server/controllers"
)

func main() {
	controllers.SetRoutes()
	http.ListenAndServe(":8888", nil)
}
