package service

import (
	"github.com/budimansol/pairproject/internal/model"
	"github.com/budimansol/pairproject/internal/repository"
)

type MenuService struct {
	repo *repository.MenuRepository
}

func NewMenuService(repo *repository.MenuRepository) *MenuService {
	return &MenuService{repo: repo}
}

func (s *MenuService) AddMenu(menu model.Menu) error {
	return s.repo.Create(menu)
}

func (s *MenuService) GetAllMenu() ([]model.Menu, error) {
	return s.repo.GetAll()
}

func (s *MenuService) UpdateMenu(menu model.Menu) error {
	return s.repo.Update(menu)
}

func (s *MenuService) DeleteMenu(id int) error {
	return s.repo.Delete(id)
}