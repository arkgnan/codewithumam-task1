package repositories

import (
	"crud-category/models"
	"database/sql"
	"errors"
	"log"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (repo *ProductRepository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	rows, err := repo.db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (repo *ProductRepository) GetProductByID(id int) (*models.Product, error) {
	var product models.Product
	err := repo.db.QueryRow("SELECT id, name, price, stock FROM products WHERE id = $1", id).Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
	if err == sql.ErrNoRows {
		return nil, errors.New("product not found")
	}
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *ProductRepository) CreateProduct(product *models.Product) error {
	log.Println(product)
	err := repo.db.QueryRow("INSERT INTO products (name, price, stock) VALUES ($1, $2, $3) RETURNING id", product.Name, product.Price, product.Stock).Scan(&product.ID)
	return err
}

func (repo *ProductRepository) UpdateProduct(product *models.Product) error {
	result, err := repo.db.Exec("UPDATE products SET name = $1, price = $2, stock = $3 WHERE id = $4", product.Name, product.Price, product.Stock, product.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("product not found")
	}
	return nil
}

func (repo *ProductRepository) DeleteProduct(id int) error {
	result, err := repo.db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("product not found")
	}
	return nil
}
