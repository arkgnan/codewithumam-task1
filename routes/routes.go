package routes

import (
	"crud-category/controller"
	"encoding/json"
	"net/http"
)

func RegisterRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "Server is up and running",
		})
	})

	router.HandleFunc("GET /api/products", controller.GetAllProducts)
	router.HandleFunc("GET /api/products/{id}", controller.GetProductByID)
	router.HandleFunc("POST /api/products", controller.StoreProduct)
	router.HandleFunc("PUT /api/products/{id}", controller.UpdateProduct)
	router.HandleFunc("DELETE /api/products/{id}", controller.DeleteProduct)

	router.HandleFunc("GET /api/categories", controller.GetAllCategories)
	router.HandleFunc("GET /api/categories/{id}", controller.GetCategoryByID)
	router.HandleFunc("POST /api/categories", controller.StoreCategory)
	router.HandleFunc("PUT /api/categories/{id}", controller.UpdateCategory)
	router.HandleFunc("DELETE /api/categories/{id}", controller.DeleteCategory)

	return router
}
