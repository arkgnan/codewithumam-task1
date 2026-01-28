package routes

import (
	"crud-category/handlers"
	"crud-category/repositories"
	"crud-category/services"
	"database/sql"
	"encoding/json"
	"net/http"
)

func RegisterRoutes(db *sql.DB) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "Server is up and running",
		})
	})

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)
	router.HandleFunc("GET /api/products", productHandler.GetAllProducts)
	router.HandleFunc("GET /api/products/{id}", productHandler.GetProductByID)
	router.HandleFunc("POST /api/products", productHandler.CreateProduct)
	router.HandleFunc("PUT /api/products/{id}", productHandler.UpdateProduct)
	router.HandleFunc("DELETE /api/products/{id}", productHandler.DeleteProduct)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	router.HandleFunc("GET /api/categories", categoryHandler.GetAllCategories)
	router.HandleFunc("GET /api/categories/{id}", categoryHandler.GetCategoryByID)
	router.HandleFunc("POST /api/categories", categoryHandler.CreateCategory)
	router.HandleFunc("PUT /api/categories/{id}", categoryHandler.UpdateCategory)
	router.HandleFunc("DELETE /api/categories/{id}", categoryHandler.DeleteCategory)

	return router
}
