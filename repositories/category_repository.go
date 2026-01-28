package repositories

import (
	"crud-category/models"
	"database/sql"
	"errors"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) CreateCategory(category *models.Category) error {
	err := r.db.QueryRow("INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id", category.Name, category.Description).Scan(&category.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepository) GetCategory(id int) (*models.Category, error) {
	row := r.db.QueryRow("SELECT id, name, description FROM categories WHERE id = $1", id)
	var category models.Category
	err := row.Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) UpdateCategory(category *models.Category) error {
	_, err := r.db.Exec("UPDATE categories SET name = $1, description = $2 WHERE id = $3", category.Name, category.Description, category.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepository) DeleteCategory(id int) error {
	_, err := r.db.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *CategoryRepository) GetAllCategories() ([]models.Category, error) {
	rows, err := r.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (r *CategoryRepository) GetCategoryByID(id int) (*models.Category, error) {
	row := r.db.QueryRow("SELECT id, name, description FROM categories WHERE id = $1", id)
	var category models.Category
	err := row.Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	return &category, nil
}
