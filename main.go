package main

import (
	"net/http"

	"curso.alura.web_app/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
