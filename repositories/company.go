package repositories

import (
	"github.com/APoniatowski/listingd/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyRepository interface {
	Create(company *models.Company) error
	Update(company *models.Company) error
	Delete(id uuid.UUID) error
	Get(id uuid.UUID) (*models.Company, error)
}

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &companyRepository{db}
}

func (r *companyRepository) Create(company *models.Company) error {
	return r.db.Create(company).Error
}

func (r *companyRepository) Update(company *models.Company) error {
	return r.db.Model(&models.Company{}).Where("id = ?", company.ID).Updates(company).Error
}

func (r *companyRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Company{}, id).Error
}

func (r *companyRepository) Get(id uuid.UUID) (*models.Company, error) {
	var company models.Company
	err := r.db.First(&company, id).Error
	if err != nil {
		return nil, err
	}
	return &company, nil
}
