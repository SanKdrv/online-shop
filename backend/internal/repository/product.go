package repository

import (
	"backend/internal/domain"
	"fmt"
	"gorm.io/gorm"
)

type ProductsRepo struct {
	db *gorm.DB
}

func NewProductsRepo(db *gorm.DB) *ProductsRepo {
	return &ProductsRepo{
		db: db,
	}
}

func (r *ProductsRepo) CreateProduct(product domain.Product) (int64, error) {
	if err := r.db.Create(&product).Error; err != nil {
		return 0, err
	}
	return product.Id, nil
}

func (r *ProductsRepo) Get(name string, brandId int64, categoryId int64) (domain.Product, error) {
	var product domain.Product
	if err := r.db.Model(&domain.Product{}).Where("name = ?, brand_id = ?, category_id = ?", name, brandId, categoryId).
		First(&product).Error; err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r *ProductsRepo) GetAllByCategoryPaginated(categoryId int64, page int, limit int) ([]domain.Product, error) {
	var products []domain.Product
	offset := (page - 1) * limit

	if err := r.db.Model(&domain.Product{}).
		Where("category_id = ?", categoryId).
		Limit(limit).
		Offset(offset).
		Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductsRepo) GetAllByNamePaginated(name string, page int, limit int) ([]domain.Product, error) {
	var products []domain.Product
	offset := (page - 1) * limit

	if err := r.db.Model(&domain.Product{}).Where("name = ?", name).
		Limit(limit).
		Offset(offset).
		Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductsRepo) GetAllByBrandPaginated(brandId int64, page int, limit int) ([]domain.Product, error) {
	var products []domain.Product
	offset := (page - 1) * limit

	if err := r.db.Model(&domain.Product{}).Where("brand_id = ?", brandId).
		Limit(limit).
		Offset(offset).
		Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductsRepo) UpdateProduct(product domain.Product) error {
	return r.db.Model(&domain.Product{}).Where("id = ?", product.Id).Updates(product).Error
}

func (r *ProductsRepo) DeleteProduct(productId int64) error {
	err := r.db.Delete(&domain.Product{Id: productId}).Error
	return err
}

func (r *ProductsRepo) GetAll(offset int64, limit int64) ([]domain.Product, int64, error) {
	var products []domain.Product
	var total int64

	// Сначала получаем общее количество
	if err := r.db.Model(&domain.Product{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count products: %w", err)
	}

	// Затем получаем rows с пагинацией
	if err := r.db.
		Offset(int(offset)).
		Limit(int(limit)).
		Find(&products).
		Error; err != nil {
		return nil, 0, fmt.Errorf("failed to get products: %w", err)
	}

	return products, total, nil
}
