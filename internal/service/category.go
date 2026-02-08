package service

import (
	"kasir-api/internal/model"
	"kasir-api/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll() ([]model.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) GetByID(id int64) (*model.Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) Create(category *model.Category) (*model.Category, error) {
	return s.repo.Create(category)
}

func (s *CategoryService) Update(id int64, category *model.Category) error {
	return s.repo.Update(id, category)
}

func (s *CategoryService) Delete(id int64) error {
	return s.repo.Delete(id)
}
