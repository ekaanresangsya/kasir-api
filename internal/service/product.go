package service

import (
	"kasir-api/internal/model"
	"kasir-api/internal/repository"
	"log"
)

type ProductService struct {
	productRepo  *repository.ProductRepository
	categoryRepo *repository.CategoryRepository
}

func NewProductService(productRepo *repository.ProductRepository, categoryRepo *repository.CategoryRepository) *ProductService {
	return &ProductService{productRepo: productRepo, categoryRepo: categoryRepo}
}

func (s *ProductService) GetAll(req model.GetProductReq) ([]model.Product, error) {
	return s.productRepo.GetAll(req)
}

func (s *ProductService) GetByID(id int) (*model.Product, error) {
	return s.productRepo.GetByID(id)
}

func (s *ProductService) Create(product *model.Product) error {

	// create product
	err := s.productRepo.Create(product)
	if err != nil {
		log.Printf("error creating product, got %v", err)
		return err
	}

	// get data category
	category, err := s.categoryRepo.GetByID(product.CategoryID)
	if err != nil {
		log.Printf("error getting category, got %v", err)
		return err
	}

	product.Category = *category

	return nil
}

func (s *ProductService) Update(product *model.Product) error {

	err := s.productRepo.Update(product)
	if err != nil {
		log.Printf("error updating product, got %v", err)
		return err
	}

	category, err := s.categoryRepo.GetByID(product.CategoryID)
	if err != nil {
		log.Printf("error getting category, got %v", err)
		return err
	}

	product.Category = *category

	return nil
}

func (s *ProductService) Delete(id int64) error {
	return s.productRepo.Delete(id)
}
