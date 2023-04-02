package services

import (
	"github.com/APoniatowski/listingd/pkg/models"
	"github.com/google/uuid"
)

type CompanyRepository interface {
	Create(company *models.Company) error
	Update(company *models.Company) error
	Delete(id uuid.UUID) error
	Get(id uuid.UUID) (*models.Company, error)
}

type CompanyService interface {
	Create(company *models.Company) error
	Update(original, updated *models.Company) error
	Delete(id uuid.UUID) error
	Get(id uuid.UUID) (*models.Company, error)
}

type CompanyServiceImpl struct {
	repo CompanyRepository
}

func NewCompanyService(repo CompanyRepository) CompanyService {
	return &CompanyServiceImpl{repo: repo}
}

func (s *CompanyServiceImpl) Create(company *models.Company) error {
	// Business logic (if any) and call to repository
	return s.repo.Create(company)
}

func (s *CompanyServiceImpl) Update(company, updated *models.Company) error {
	// Business logic (if any) and call to repository
	return s.repo.Update(updated)
}

func (s *CompanyServiceImpl) Delete(id uuid.UUID) error {
	// Business logic (if any) and call to repository
	return s.repo.Delete(id)
}

func (s *CompanyServiceImpl) Get(id uuid.UUID) (*models.Company, error) {
	// Business logic (if any) and call to repository
	return s.repo.Get(id)
}
