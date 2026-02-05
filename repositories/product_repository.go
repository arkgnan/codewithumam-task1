package repositories

import (
	"crud-category/models"
	"database/sql"
	"errors"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (repo *ProductRepository) GetProducts(name string) ([]models.Product, error) {
	var products []models.Product
	var categoryID sql.NullInt64    // Use sql.NullInt64 for nullable integer columns
	var categoryName sql.NullString // Use sql.NullString for nullable string columns
	var categoryDescription sql.NullString
	query := `
		SELECT
			p.id, p.name, p.price, p.stock,
			c.id, c.name, c.description
		FROM
			products p
		LEFT JOIN
			categories c ON p.category_id = c.id
	`
	args := []any{}
	if name != "" {
		query += " WHERE p.name ILIKE $1"
		args = append(args, "%"+name+"%")
	}
	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(
			&product.ID, &product.Name, &product.Price, &product.Stock,
			&categoryID, &categoryName, &categoryDescription,
		); err != nil {
			return nil, err
		}
		if categoryID.Valid {
			categoryInt := int(categoryID.Int64)
			product.Category = &models.Category{
				ID:          categoryInt,
				Name:        categoryName.String,
				Description: categoryDescription.String,
			}
		} else {
			// If categoryID is NULL, set Category to nil
			product.Category = nil
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
	var categoryID sql.NullInt64    // Use sql.NullInt64 for nullable integer columns
	var categoryName sql.NullString // Use sql.NullString for nullable string columns
	var categoryDescription sql.NullString
	query := `
		SELECT
			p.id, p.name, p.price, p.stock,
			c.id, c.name, c.description
		FROM
			products p
		LEFT JOIN
			categories c ON p.category_id = c.id
		WHERE
			p.id = $1
	`
	err := repo.db.QueryRow(query, id).Scan(
		&product.ID, &product.Name, &product.Price, &product.Stock,
		&categoryID, &categoryName, &categoryDescription,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("product not found")
	}
	if err != nil {
		return nil, err
	}

	if categoryID.Valid {
		categoryInt := int(categoryID.Int64)
		product.Category = &models.Category{
			ID:          categoryInt,
			Name:        categoryName.String,
			Description: categoryDescription.String,
		}
	} else {
		// If categoryID is NULL, set Category to nil
		product.Category = nil
	}
	return &product, nil
}

func (repo *ProductRepository) CreateProduct(product *models.ProductRequest) error {
	err := repo.db.QueryRow("INSERT INTO products (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id", product.Name, product.Price, product.Stock, product.CategoryID).Scan(&product.ID)
	return err
}

func (repo *ProductRepository) UpdateProduct(product *models.ProductRequest) error {
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
