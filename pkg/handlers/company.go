package handlers

import (
	"net/http"

	"github.com/APoniatowski/listingd/pkg/models"
	"github.com/APoniatowski/listingd/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CompanyHandler struct {
	companyService services.CompanyService
}

func NewCompanyHandler(companyService services.CompanyService) *CompanyHandler {
	return &CompanyHandler{
		companyService: companyService,
	}
}

func (h *CompanyHandler) CreateCompany(c *gin.Context) {
	var input models.Company
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company := &models.Company{
		ID:                uuid.New(),
		Name:              input.Name,
		Description:       input.Description,
		AmountOfEmployees: input.AmountOfEmployees,
		Registered:        input.Registered,
		Type:              input.Type,
	}

	if err := h.companyService.Create(company); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, company)
}

func (h *CompanyHandler) PatchCompany(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	company, err := h.companyService.Get(id)
	if err != nil {
		if err == models.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var input models.Company
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.companyService.Update(company, &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, company)
}

func (h *CompanyHandler) DeleteCompany(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.companyService.Delete(id); err != nil {
		if err == models.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *CompanyHandler) GetCompany(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	company, err := h.companyService.Get(id)
	if err != nil {
		if err == models.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, company)
}
