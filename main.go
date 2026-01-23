package main

import (
	"crud-category/controller"
	"encoding/json"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	router.HandleFunc("GET /api/products", controller.GetAllProducts)
	router.HandleFunc("GET /api/products/{id}", controller.GetProductByID)
	router.HandleFunc("POST /api/products", controller.StoreProduct)
	router.HandleFunc("PUT /api/products/{id}", controller.UpdateProduct)
	router.HandleFunc("DELETE /api/products/{id}", controller.DeleteProduct)

	println("Server started on port 8080")
	http.ListenAndServe(":8080", router)
}
