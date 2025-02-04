package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type CategoriesRepo struct {
	db *gorm.DB
}

func NewCategoriesRepo(db *gorm.DB) *CategoriesRepo {
	return &CategoriesRepo{
		db: db,
	}
}

func (r *CategoriesRepo) GetIdByCategory(name string) (int64, error) {
	var id int64
	err := r.db.Model(&domain.Category{}).Where("name = ?", name).Pluck("id", &id).Error
	return id, err
}

func (r *CategoriesRepo) GetCategoryById(categoryId int64) (string, error) {
	var name string
	err := r.db.Model(&domain.Category{}).Where("id = ?", categoryId).Pluck("name", &name).Error
	return name, err
}

func (r *CategoriesRepo) CreateCategory(name string) (int64, error) {
	var category = domain.Category{
		Name: name,
	}
	if err := r.db.Create(&category).Error; err != nil {
		return 0, err
	}
	return category.ID, nil
}

func (r *CategoriesRepo) DeleteCategory(categoryID int64) error {
	if err := r.db.Delete(&domain.Category{ID: categoryID}).Error; err != nil {
		return err
	}
	return nil
}

func (r *CategoriesRepo) UpdateCategory(categoryID int64, name string) error {
	err := r.db.Model(&domain.Category{}).Where("id = ?", categoryID).Update("name", name).Error
	return err
}
