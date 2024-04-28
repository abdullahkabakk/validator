package validator

import (
	"github.com/abdullahkabakk/validator/internal/validator/locales"
	"reflect"
	"testing"
)

// TestValidateSpecialCharacter tests the validateSpecialCharacter function.
func TestValidateSpecialCharacter(t *testing.T) {
	// Define test cases
	tests := []struct {
		name        string // Test case name
		value       string // Input value
		expectedErr bool   // Expected error presence
	}{
		{name: "WithSpecialCharacter", value: "hello$world", expectedErr: false},
		{name: "WithoutSpecialCharacter", value: "helloworld", expectedErr: true},
		{name: "EmptyString", value: "", expectedErr: true},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up locale messages
			messages := locales.ErrorMessages{"specialCharacter": "Field %s must contain special characters"}

			// Convert value to reflect value
			value := reflect.ValueOf(tt.value)

			// Call validateSpecialCharacter
			err := validateSpecialCharacter(value, messages, "fieldName", "rule")

			// Check error
			if (err != nil) != tt.expectedErr {
				t.Errorf("Test case %s: expected error %v, got error %v", tt.name, tt.expectedErr, err)
			}
		})
	}
}

// TestContainsSpecialCharacter tests the containsSpecialCharacter function.
func TestContainsSpecialCharacter(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string // Test case name
		value    string // Input value
		expected bool   // Expected result
	}{
		{name: "WithSpecialCharacter", value: "hello$world", expected: true},
		{name: "WithoutSpecialCharacter", value: "helloworld", expected: false},
		{name: "EmptyString", value: "", expected: false},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call containsSpecialCharacter
			result := containsSpecialCharacter(tt.value)

			// Check result
			if result != tt.expected {
				t.Errorf("Test case %s: expected %v, got %v", tt.name, tt.expected, result)
			}
		})
	}
}
