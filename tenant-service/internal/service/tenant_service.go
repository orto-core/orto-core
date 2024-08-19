package service

import (
	"github.com/orto-core/server/tenant-service/internal/models"
	"github.com/orto-core/server/tenant-service/internal/repository"
)

type TenantService interface {
	AddTenant(*models.Tenant) (string, error)
}

type tenantService struct {
	repository repository.TenantRepository
}

func NewTenantService(repository repository.TenantRepository) TenantService {
	return &tenantService{
		repository: repository,
	}
}

func (s *tenantService) AddTenant(tenant *models.Tenant) (string, error) {
	return "", nil
}
