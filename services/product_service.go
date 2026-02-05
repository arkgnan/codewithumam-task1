package services

import (
	"crud-category/models"
	"crud-category/repositories"
	"errors"
)

type ProductService struct {
	repo         *repositories.ProductRepository
	categoryRepo *repositories.CategoryRepository
}

func NewProductService(repo *repositories.ProductRepository, categoryRepo *repositories.CategoryRepository) *ProductService {
	return &ProductService{
		repo:         repo,
		categoryRepo: categoryRepo,
	}
}

func (s *ProductService) GetAllProducts(name string) ([]models.Product, error) {
	return s.repo.GetProducts(name)
}

func (s *ProductService) GetProductByID(id int) (*models.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s *ProductService) CreateProduct(product *models.ProductRequest) error {
	// validate category
	if product.CategoryID != nil {
		validCategory, err := s.categoryRepo.GetCategoryByID(*product.CategoryID)
		if err != nil {
			return err
		}
		if validCategory == nil {
			return errors.New("category not found")
		}
	}

	return s.repo.CreateProduct(product)
}

func (s *ProductService) UpdateProduct(product *models.ProductRequest) error {
	if product.CategoryID != nil {
		validCategory, err := s.categoryRepo.GetCategoryByID(*product.CategoryID)
		if err != nil {
			return err
		}
		if validCategory == nil {
			return errors.New("category not found")
		}
	}
	return s.repo.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.repo.DeleteProduct(id)
}
