package service

import (
	"github.com/first-restapi-golang/internal/model"
	"github.com/first-restapi-golang/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func New(r *repository.Repository) (s *Service) {
	s = &Service{r}
	return
}

func (s *Service) Add(c *model.Category) error {
	return s.repo.AddCategoryByModel(c)
}
