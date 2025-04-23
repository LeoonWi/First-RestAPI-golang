package repository

import (
	"github.com/first-restapi-golang/internal/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) (r *Repository) {
	r = &Repository{db}
	return
}

func (r *Repository) AddCategoryByModel(c *model.Category) error {
	return r.db.Create(c).Error
}
