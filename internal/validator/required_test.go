package validator

import (
	"github.com/abdullahkabakk/validator/internal/validator/locales"
	"reflect"
	"testing"
)

// TestValidateRequired tests the validateRequired function.
func TestValidateRequired(t *testing.T) {
	// Define test cases
	tests := []struct {
		name        string      // Test case name
		value       interface{} // Input value
		expectedErr bool        // Expected error presence
	}{
		{name: "NonEmptyString", value: "hello", expectedErr: false},
		{name: "EmptyString", value: "", expectedErr: true},
		{name: "NonEmptySlice", value: []int{1, 2, 3}, expectedErr: false},
		{name: "EmptySlice", value: []int{}, expectedErr: true},
		{name: "NonEmptyMap", value: map[string]int{"a": 1}, expectedErr: false},
		{name: "EmptyMap", value: map[string]int{}, expectedErr: true},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up locale messages
			messages := locales.ErrorMessages{"required": "Field %s is required"}

			// Convert value to reflect value
			value := reflect.ValueOf(tt.value)

			// Call validateRequired
			err := validateRequired(value, messages, "fieldName", "rule")

			// Check error
			if (err != nil) != tt.expectedErr {
				t.Errorf("Test case %s: expected error %v, got error %v", tt.name, tt.expectedErr, err)
			}
		})
	}
}

// TestIsEmpty tests the isEmpty function.
func TestIsEmpty(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string      // Test case name
		value    interface{} // Input value
		expected bool        // Expected result
	}{
		{name: "EmptyString", value: "", expected: true},
		{name: "NonEmptyString", value: "hello", expected: false},
		{name: "EmptySlice", value: []int{}, expected: true},
		{name: "NonEmptySlice", value: []int{1, 2, 3}, expected: false},
		{name: "EmptyMap", value: map[string]int{}, expected: true},
		{name: "NonEmptyMap", value: map[string]int{"a": 1}, expected: false},
		// Default case
		{name: "DefaultCase", value: 0, expected: true},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert value to reflect value
			value := reflect.ValueOf(tt.value)

			// Call isEmpty
			result := isEmpty(value)

			// Check result
			if result != tt.expected {
				t.Errorf("Test case %s: expected %v, got %v", tt.name, tt.expected, result)
			}
		})
	}
}
