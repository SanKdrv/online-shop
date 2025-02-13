package repository

import (
	"backend/internal/domain"
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

func (r *ProductsRepo) GetAllByCategory(categoryId int64) ([]domain.Product, error) {
	var products []domain.Product
	if err := r.db.Model(&domain.Product{}).Where("category_id = ?", categoryId).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductsRepo) GetAllByName(name string) ([]domain.Product, error) {
	var products []domain.Product
	if err := r.db.Model(&domain.Product{}).Where("name = ?", name).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductsRepo) GetAllByBrand(brandId int64) ([]domain.Product, error) {
	var products []domain.Product
	if err := r.db.Model(&domain.Product{}).Where("brand_id = ?", brandId).Find(&products).Error; err != nil {
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
