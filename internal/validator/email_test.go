package validator

import (
	"fmt"
	"github.com/abdullahkabakk/validator/internal/validator/locales"
	"reflect"
	"testing"
)

// TestValidateEmail tests the validateEmail function.
func TestValidateEmail(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string // Test case name
		email    string // Input email string
		expected error  // Expected error
	}{
		{
			name:     "Valid Email",
			email:    "test@example.com",
			expected: nil,
		},
		{
			name:     "Empty Email",
			email:    "",                                                     // Empty email address
			expected: fmt.Errorf("Error message for empty email testField."), // Expected error message for empty email
		},
		{
			name: "Long Email",
			email: "veryultralongemailaddresslongerthan255charactersveryultralongveryultralongemailaddresslonger" +
				"than255charactersveryultralongemailaddresslongerthan255charactersveryultralongemailaddresslongerthan255" +
				"charactersveryultralongemailaddresslongerthan255charactersveryultralongemailaddresslongerthan255characters" +
				"veryultralongemailaddresslongerthan255charactersveryultralongemailaddresslongerthan255charactersveryultralong" +
				"emailaddresslongerthan255charactersveryultralongemailaddresslongerthan255charactersveryultralong" +
				"emailaddresslongerthan255characters@verylongemail.com",
			expected: fmt.Errorf("Email address for testField is too long."), // Expected error message for long email
		},
		{
			name:     "Invalid Email",
			email:    "invalidemail",                                     // Invalid email address format
			expected: fmt.Errorf("Invalid email address for testField."), // Expected error message for invalid email
		},
	}

	// Mock error messages for localization
	errorMessages := locales.ErrorMessages{
		"emailTooLong": "Email address for %s is too long.",
		"invalidEmail": "Invalid email address for %s.",
		"emailIsEmpty": "Error message for empty email %s.", // Error message for empty email
	}

	// Iterate over test cases
	for _, test := range tests {
		// Run each test case
		t.Run(test.name, func(t *testing.T) {
			// Convert email string to reflect.Value for testing
			value := reflect.ValueOf(test.email)
			// Validate the email
			err := validateEmail(value, errorMessages, "testField", "email")
			// Compare the actual error with the expected error
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
