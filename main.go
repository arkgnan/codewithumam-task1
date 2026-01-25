package main

import (
	"crud-category/middleware"
	"crud-category/routes"
	"net/http"
)

func main() {
	router := routes.RegisterRoutes()

	loggedRouter := middleware.LoggingMiddleware(router)

	println("Server started on port 8080")
	http.ListenAndServe(":8080", loggedRouter)
}
