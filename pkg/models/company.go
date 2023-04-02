package models

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyType string

const (
	Corporations       CompanyType = "Corporations"
	NonProfit          CompanyType = "NonProfit"
	Cooperative        CompanyType = "Cooperative"
	SoleProprietorship CompanyType = "Sole Proprietorship"
)

type Company struct {
	ID                uuid.UUID   `gorm:"type:uuid;primary_key" json:"id"`
	Name              string      `gorm:"type:varchar(15);unique" json:"name" binding:"required"`
	Description       string      `gorm:"type:text" json:"description"`
	AmountOfEmployees int         `json:"amount_of_employees" binding:"required"`
	Registered        bool        `json:"registered" binding:"required"`
	Type              CompanyType `json:"type" binding:"required,oneof=Corporations NonProfit Cooperative Sole Proprietorship"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	DeletedAt         *time.Time  `gorm:"index" json:"deleted_at"`
}

var ErrNotFound = errors.New("company not found")

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (c *Company) Validate() error {
	return validate.Struct(c)
}

func (company *Company) BeforeCreate(tx *gorm.DB) (err error) {
	company.ID = uuid.New()
	return
}
