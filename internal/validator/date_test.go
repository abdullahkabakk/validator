package validator

import (
	"fmt"
	"github.com/abdullahkabakk/validator/internal/validator/locales"
	"reflect"
	"testing"
)

// TestValidateDate tests the validateDate function.
func TestValidateDate(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string // Test case name
		date     string // Input date string
		expected error  // Expected error
	}{
		{
			name:     "Valid Date",
			date:     "2024-04-27",
			expected: nil,
		},
		{
			name:     "Invalid Date",
			date:     "2024-02-30",                              // February 30th, an invalid date
			expected: fmt.Errorf("Invalid date for testField."), // Expected error message for invalid date
		},
		{
			name:     "Date Too Long",
			date:     "2024-04-27T12:00:00",                         // Date with time, which exceeds the maximum length
			expected: fmt.Errorf("Date for testField is too long."), // Expected error message for date too long
		},
	}

	// Define mock error messages for localization
	errorMessages := locales.ErrorMessages{
		"dateTooLong": "Date for %s is too long.",
		"invalidDate": "Invalid date for %s.",
	}

	// Iterate over test cases
	for _, test := range tests {
		// Run each test case
		t.Run(test.name, func(t *testing.T) {
			// Convert date string to reflect.Value for testing
			value := reflect.ValueOf(test.date)
			// Validate the date
			err := validateDate(value, errorMessages, "testField", "date")
			// Check if the error matches the expected error
			if err == nil && test.expected != nil {
				t.Errorf("Expected error, got nil")
			} else if err != nil && test.expected == nil {
				t.Errorf("Expected no error, got %v", err)
			} else if err != nil && test.expected != nil && err.Error() != test.expected.Error() {
				t.Errorf("Expected error message '%s', got '%s'", test.expected.Error(), err.Error())
			}
		})
	}
}
