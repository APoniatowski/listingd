package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/APoniatowski/listingd/pkg/models"
)

func TestGetCompany(t *testing.T) {
	// Create a new company for testing
	company := models.Company{
		Name:              "Test Company",
		Description:       "This is a test company",
		AmountOfEmployees: 50,
		Registered:        true,
		Type:              models.Corporations,
	}
	err := db.Create(&company).Error
	if err != nil {
		t.Fatalf("Error creating test company: %v", err)
	}

	// Prepare the test request
	req, _ := http.NewRequest("GET", fmt.Sprintf("/companies/%s", company.ID.String()), nil)

	// Execute the request using the router from main_test.go
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response status
	if w.Code != http.StatusOK {
		t.Fatalf("Expected HTTP status 200, got %d", w.Code)
	}

	// Parse the response JSON
	var respCompany models.Company
	if err := json.Unmarshal(w.Body.Bytes(), &respCompany); err != nil {
		t.Fatalf("Error unmarshalling response JSON: %v", err)
	}

	// Check if the response contains the correct company data
	if respCompany.ID != company.ID {
		t.Errorf("Expected ID to be %s, got %s", company.ID, respCompany.ID)
	}
	if respCompany.Name != company.Name {
		t.Errorf("Expected Name to be %s, got %s", company.Name, respCompany.Name)
	}
	if respCompany.Description != company.Description {
		t.Errorf("Expected Description to be %s, got %s", company.Description, respCompany.Description)
	}
	if respCompany.AmountOfEmployees != company.AmountOfEmployees {
		t.Errorf("Expected AmountOfEmployees to be %d, got %d", company.AmountOfEmployees, respCompany.AmountOfEmployees)
	}
	if respCompany.Registered != company.Registered {
		t.Errorf("Expected Registered to be %v, got %v", company.Registered, respCompany.Registered)
	}
	if respCompany.Type != company.Type {
		t.Errorf("Expected Type to be %s, got %s", company.Type, respCompany.Type)
	}
}
