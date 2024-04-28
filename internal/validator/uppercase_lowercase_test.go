package validator

import (
	"github.com/abdullahkabakk/validator/internal/validator/locales"
	"reflect"
	"testing"
)

// TestValidateUppercase tests the validateUppercase function.
func TestValidateUppercase(t *testing.T) {
	// Define test cases
	tests := []struct {
		name        string // Test case name
		value       string // Input value
		expectedErr bool   // Expected error presence
	}{
		{name: "WithUppercase", value: "HelloWorld", expectedErr: false},
		{name: "WithoutUppercase", value: "helloworld", expectedErr: true},
		{name: "EmptyString", value: "", expectedErr: true},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up locale messages
			messages := locales.ErrorMessages{"uppercaseLetter": "Field %s must contain at least one uppercase letter"}

			// Convert value to reflect value
			value := reflect.ValueOf(tt.value)

			// Call validateUppercase
			err := validateUppercase(value, messages, "fieldName", "rule")

			// Check error
			if (err != nil) != tt.expectedErr {
				t.Errorf("Test case %s: expected error %v, got error %v", tt.name, tt.expectedErr, err)
			}
		})
	}
}

// TestValidateLowercase tests the validateLowercase function.
func TestValidateLowercase(t *testing.T) {
	// Define test cases
	tests := []struct {
		name        string // Test case name
		value       string // Input value
		expectedErr bool   // Expected error presence
	}{
		{name: "WithLowercase", value: "helloWorld", expectedErr: false},
		{name: "WithoutLowercase", value: "HELLOWORLD", expectedErr: true},
		{name: "EmptyString", value: "", expectedErr: true},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up locale messages
			messages := locales.ErrorMessages{"lowercaseLetter": "Field %s must contain at least one lowercase letter"}

			// Convert value to reflect value
			value := reflect.ValueOf(tt.value)

			// Call validateLowercase
			err := validateLowercase(value, messages, "fieldName", "rule")

			// Check error
			if (err != nil) != tt.expectedErr {
				t.Errorf("Test case %s: expected error %v, got error %v", tt.name, tt.expectedErr, err)
			}
		})
	}
}

// TestContainsUppercase tests the containsUppercase function.
func TestContainsUppercase(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string // Test case name
		value    string // Input value
		expected bool   // Expected result
	}{
		{name: "WithUppercase", value: "HelloWorld", expected: true},
		{name: "WithoutUppercase", value: "helloworld", expected: false},
		{name: "EmptyString", value: "", expected: false},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call containsUppercase
			result := containsUppercase(tt.value)

			// Check result
			if result != tt.expected {
				t.Errorf("Test case %s: expected %v, got %v", tt.name, tt.expected, result)
			}
		})
	}
}

// TestContainsLowercase tests the containsLowercase function.
func TestContainsLowercase(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string // Test case name
		value    string // Input value
		expected bool   // Expected result
	}{
		{name: "WithLowercase", value: "helloWorld", expected: true},
		{name: "WithoutLowercase", value: "HELLOWORLD", expected: false},
		{name: "EmptyString", value: "", expected: false},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call containsLowercase
			result := containsLowercase(tt.value)

			// Check result
			if result != tt.expected {
				t.Errorf("Test case %s: expected %v, got %v", tt.name, tt.expected, result)
			}
		})
	}
}
