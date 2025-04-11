package repository

import (
	"backend/internal/domain"
	"fmt"
	"gorm.io/gorm"
)

type BrandsRepo struct {
	db *gorm.DB
}

func NewBrandsRepo(db *gorm.DB) *BrandsRepo {
	return &BrandsRepo{
		db: db,
	}
}

func (r *BrandsRepo) GetIdByBrand(name string) (int64, error) {
	var id int64
	err := r.db.Model(&domain.Brand{}).Where("name = ?", name).Pluck("id", &id).Error
	return id, err
}

func (r *BrandsRepo) GetBrandById(brandId int64) (string, error) {
	var name string
	err := r.db.Model(&domain.Brand{}).Where("id = ?", brandId).Pluck("name", &name).Error
	return name, err
}

func (r *BrandsRepo) CreateBrand(name string) (int64, error) {
	var brand = domain.Brand{
		Name: name,
	}
	if err := r.db.Create(&brand).Error; err != nil {
		return 0, err
	}
	return brand.Id, nil
}

func (r *BrandsRepo) DeleteBrand(brandId int64) error {
	var brand = domain.Brand{
		Id: brandId,
	}
	// Ищем запись в базе данных
	if err := r.db.Where("id = ?", brand.Id).First(&brand).Error; err != nil {
		return err
	}

	// Удаляем найденный кортеж
	if err := r.db.Delete(&brand).Error; err != nil {
		return err
	}

	// Возвращаем найденную запись
	return nil
}

func (r *BrandsRepo) UpdateBrand(brandId int64, name string) error {
	// Обновляем только поле Name у нужного бренда
	if err := r.db.Model(&domain.Brand{}).
		Where("id = ?", brandId).
		Update("name", name).Error; err != nil {
		return err
	}

	return nil
}

func (r *BrandsRepo) GetAll(offset int64, limit int64) ([]domain.Brand, int64, error) {
	var brands []domain.Brand
	var total int64

	// Сначала получаем общее количество брендов
	if err := r.db.Model(&domain.Brand{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count brands: %w", err)
	}

	// Затем получаем бренды с пагинацией
	if err := r.db.
		Offset(int(offset)).
		Limit(int(limit)).
		Find(&brands).
		Error; err != nil {
		return nil, 0, fmt.Errorf("failed to get brands: %w", err)
	}

	return brands, total, nil
}
