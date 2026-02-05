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
	categoryRepo := repositories.NewCategoryRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)

	productService := services.NewProductService(productRepo, categoryRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	transactionService := services.NewTransactionService(transactionRepo)

	productHandler := handlers.NewProductHandler(productService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	router.HandleFunc("GET /api/products", productHandler.GetAllProducts)
	router.HandleFunc("GET /api/products/{id}", productHandler.GetProductByID)
	router.HandleFunc("POST /api/products", productHandler.CreateProduct)
	router.HandleFunc("PUT /api/products/{id}", productHandler.UpdateProduct)
	router.HandleFunc("DELETE /api/products/{id}", productHandler.DeleteProduct)

	router.HandleFunc("GET /api/categories", categoryHandler.GetAllCategories)
	router.HandleFunc("GET /api/categories/{id}", categoryHandler.GetCategoryByID)
	router.HandleFunc("POST /api/categories", categoryHandler.CreateCategory)
	router.HandleFunc("PUT /api/categories/{id}", categoryHandler.UpdateCategory)
	router.HandleFunc("DELETE /api/categories/{id}", categoryHandler.DeleteCategory)

	router.HandleFunc("POST /api/checkout", transactionHandler.Checkout)
	router.HandleFunc("GET /api/reports/transaction", transactionHandler.TransactionReport)

	return router
}
