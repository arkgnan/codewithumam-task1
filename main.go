package main

import (
	"crud-category/routes"
	"net/http"
)

func main() {
	router := routes.RegisterRoutes()

	println("Server started on port 8080")
	http.ListenAndServe(":8080", router)
}
