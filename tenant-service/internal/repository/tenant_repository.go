package repository

import (
	"github.com/orto-core/server/tenant-service/internal/models"
	"github.com/orto-core/server/tenant-service/internal/store"
	"gorm.io/gorm"
)

type TenantRepository interface {
	CreateTenant(*models.Tenant) error
	GetTenantById(uint) (models.Tenant, error)
	GetTenants() ([]models.Tenant, error)
	UpdateTenant(*models.Tenant) error
	DeleteTenantById(uint) error
}

type tenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) TenantRepository {
	return &tenantRepository{
		db: store.DB,
	}
}

func (r *tenantRepository) CreateTenant(tenant *models.Tenant) error {
	if err := r.db.Create(&tenant).Error; err != nil {
		return err
	}
	return nil
}

func (r *tenantRepository) GetTenantById(id uint) (models.Tenant, error) {
	var tenant models.Tenant
	if err := r.db.First(&tenant, id).Error; err != nil {
		return models.Tenant{}, err
	}
	return tenant, nil
}

func (r *tenantRepository) GetTenants() ([]models.Tenant, error) {
	var tenants []models.Tenant
	if err := r.db.Find(&tenants).Error; err != nil {
		return nil, err
	}
	return tenants, nil
}

func (r *tenantRepository) UpdateTenant(tenant *models.Tenant) error {
	if err := r.db.Model(&tenant).Updates(models.Tenant{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *tenantRepository) DeleteTenantById(id uint) error {
	if err := r.db.Delete(&models.Tenant{}).Error; err != nil {
		return err
	}
	return nil
}
