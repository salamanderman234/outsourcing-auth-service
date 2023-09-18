package repository

import (
	"context"

	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	model "github.com/salamanderman234/outsourcing-auth-profile-service/models"
	"gorm.io/gorm"
)

type partnerRepository struct {
	client *gorm.DB
}

func NewPartnerRepository(client *gorm.DB) domain.PartnerRepository {
	return &partnerRepository {
		client: client,
	}
}

func(r partnerRepository) Create(ctx context.Context, data model.Partner) (error)  {
	result := r.client.WithContext(ctx).Create(&data)
	if result != nil {
		return result.Error
	}
	return nil
}
func(r partnerRepository) Get(ctx context.Context, filter model.Partner) ([]model.Partner, error) {
	partners  := []model.Partner{}
	result := r.client.WithContext(ctx).Model(&filter).Where(&filter).Find(&partners) 
	if result.Error != nil {
		return nil, result.Error
	}
	return partners, nil
}
func(r partnerRepository) FindById(ctx context.Context, id uint) (model.Partner, error) {
	data := model.Partner{}
	result := r.client.WithContext(ctx).Where("id = ?", id).First(&data)
	if result != nil {
		return data, result.Error
	}
	return data,nil
}
func(r partnerRepository) Update(ctx context.Context, id uint, data model.Partner) (int, error) {
	result := r.client.WithContext(ctx).
		Model(&model.Partner{}).
		Where("id = ?", id).
		Updates(&data)
	if result.Error != nil {
		return 0, nil
	}
	return int(result.RowsAffected), nil
}

func(r partnerRepository) Delete(ctx context.Context, id uint) (int, error) {
	result := r.client.WithContext(ctx).
		Where("id = ?", id).
		Delete(&model.Partner{})
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), nil
}