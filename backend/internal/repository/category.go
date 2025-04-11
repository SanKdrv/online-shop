package repository

import (
	"backend/internal/domain"
	"fmt"
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
	return category.Id, nil
}

func (r *CategoriesRepo) DeleteCategory(categoryId int64) error {
	if err := r.db.Delete(&domain.Category{Id: categoryId}).Error; err != nil {
		return err
	}
	return nil
}

func (r *CategoriesRepo) UpdateCategory(categoryId int64, name string) error {
	err := r.db.Model(&domain.Category{}).Where("id = ?", categoryId).Update("name", name).Error
	return err
}

func (r *CategoriesRepo) GetAll(offset int64, limit int64) ([]domain.Category, int64, error) {
	var categories []domain.Category
	var total int64

	// Сначала получаем общее количество
	if err := r.db.Model(&domain.Category{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count categories: %w", err)
	}

	// Затем получаем rows с пагинацией
	if err := r.db.
		Offset(int(offset)).
		Limit(int(limit)).
		Find(&categories).
		Error; err != nil {
		return nil, 0, fmt.Errorf("failed to get categories: %w", err)
	}

	return categories, total, nil
}
