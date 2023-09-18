package repository

import (
	"context"

	domain "github.com/salamanderman234/outsourcing-auth-profile-service/domains"
	model "github.com/salamanderman234/outsourcing-auth-profile-service/models"
	"gorm.io/gorm"
)

type adminRepository struct {
	client *gorm.DB
}

func NewAdminRepository(client *gorm.DB) domain.AdminRepository {
	return &adminRepository{
		client: client,
	}
}

func(r adminRepository) Create(ctx context.Context, data model.Admin) (error)  {
	result := r.client.WithContext(ctx).Create(&data)
	return result.Error
}
func(r adminRepository) Get(ctx context.Context, filter model.Admin) ([]model.Admin, error) {
	admins := []model.Admin{}
	result := r.client.WithContext(ctx).Model(&filter).Where(&filter).Find(&admins) 
	return admins, result.Error
}
func(r adminRepository) FindById(ctx context.Context, id uint) (model.Admin, error) {
	data := model.Admin{}
	result := r.client.WithContext(ctx).Where("id = ?", id).First(&data)
	return data,result.Error
}
func(r adminRepository) Update(ctx context.Context, id uint, data model.Admin) (int, error) {
	result := r.client.WithContext(ctx).
		Model(&model.Admin{}).
		Where("id = ?", id).
		Updates(&data)
	return int(result.RowsAffected), result.Error
}

func(r adminRepository) Delete(ctx context.Context, id uint) (int, error) {
	result := r.client.WithContext(ctx).
		Where("id = ?", id).
		Delete(&model.Admin{})
	return int(result.RowsAffected), result.Error
}