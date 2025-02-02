package repository

import "gorm.io/gorm"

type CategoriesRepo struct {
	db *gorm.DB
}

func NewCategoriesRepo(db *gorm.DB) *CategoriesRepo {
	return &CategoriesRepo{
		db: db,
	}
}

func (r *CategoriesRepo) GetIDByCategory(name string) (int64, error) {
	return 0, nil
}

func (r *CategoriesRepo) GetCategoryByID(categoryID int64) (string, error) {
	return "", nil
}

func (r *CategoriesRepo) CreateCategory(name string) (int64, error) {
	return 0, nil
}

func (r *CategoriesRepo) DeleteCategory(categoryID int64) error {
	return nil
}

func (r *CategoriesRepo) UpdateCategory(categoryID int64, name string) error {
	return nil
}
