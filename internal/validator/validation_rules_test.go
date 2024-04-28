package validator

import (
	"errors"
	"testing"
)

type User struct {
	Username string `validate:"required,en=Username"`
	Password string `validate:"required,min=8,max=20"`
}

// TestValidateStruct tests the validation of a struct using the validateStruct function.
func TestValidateStruct(t *testing.T) {
	// Register default validation rules
	RegisterDefaultValidationRules()

	// Define test cases
	testCases := []struct {
		name   string      // Name of the test case
		input  interface{} // Input struct to be validated
		lang   string      // Language for error messages
		expect error       // Expected error (nil if no error expected)
	}{
		{
			name: "ValidUser", // Test case for a valid user
			input: User{ // Input user struct
				Username: "testuser",
				Password: "securepassword",
			},
			lang:   "en", // Language for error messages
			expect: nil,  // Expected error (nil for valid input)
		},
		{
			name: "InvalidUser", // Test case for an invalid user
			input: User{ // Input user struct with invalid password
				Username: "test",
				Password: "weak",
			},
			lang:   "en",                                                      // Language for error messages
			expect: errors.New("Password must be at least 8 characters long"), // Expected error message
		},
		// input: nil
		{
			name:   "NilInput",                 // Test case for nil input
			input:  nil,                        // Nil input
			lang:   "en",                       // Language for error messages
			expect: errors.New("input is nil"), // Expected error message
		},
		// input: not a struct
		{
			name:   "NonStructInput",                    // Test case for non-struct input
			input:  "test",                              // Non-struct input
			lang:   "en",                                // Language for error messages
			expect: errors.New("input is not a struct"), // Expected error message
		},
		// invalid language
		{
			name: "InvalidLanguage", // Test case for invalid language
			input: User{ // Input user struct
				Username: "testuser",
				Password: "securepassword",
			},
			lang:   "",  // Invalid empty language
			expect: nil, // No error expected for invalid language
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call validateStruct and check for errors
			err := ValidateStruct(tc.input, tc.lang)
			if (err != nil && tc.expect == nil) || (err == nil && tc.expect != nil) || (err != nil && tc.expect != nil && err.Error() != tc.expect.Error()) {
				t.Errorf("Test case %s failed: expected error '%v', got '%v'", tc.name, tc.expect, err)
			}
		})
	}
}
