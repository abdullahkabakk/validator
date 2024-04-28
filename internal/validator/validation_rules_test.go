package validator

import (
	"errors"
	"github.com/abdullahkabakk/validator/internal/validator/locales"
	"sync"
	"testing"
)

type User struct {
	Username string `validate:"required,en=Username"`
	Password string `validate:"required,min=8,max=20"`
}

// TestValidateStruct tests the validation of a struct using the validateStruct function.
func TestValidateStruct(t *testing.T) {
	// Register default validation rules
	registerDefaultValidationRules()

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
			err := validateStruct(tc.input, tc.lang)
			if (err != nil && tc.expect == nil) || (err == nil && tc.expect != nil) || (err != nil && tc.expect != nil && err.Error() != tc.expect.Error()) {
				t.Errorf("Test case %s failed: expected error '%v', got '%v'", tc.name, tc.expect, err)
			}
		})
	}
}

// BenchmarkValidationWithDifferentLanguages benchmarks the validation process with different languages.
func BenchmarkValidationWithDifferentLanguages(b *testing.B) {
	// Specify the languages for which we will load messages
	languages := []string{"en", "tr"} // For example, English and Turkish

	// Load the messages for each language and store them in a map
	messages := make(map[string]locales.ErrorMessages)
	var wg sync.WaitGroup

	// Increment the WaitGroup counter for each language
	for _, lang := range languages {
		wg.Add(1)
		go func(lang string) {
			defer wg.Done()

			msg, err := locales.LoadMessagesFromJSON(lang)
			if err != nil {
				b.Fatalf("Error loading messages for language %s: %v", lang, err)
			}
			messages[lang] = msg
		}(lang)
	}

	// Wait for all language loading goroutines to complete
	wg.Wait()

	// Create our validator and initialize it with language-specific validation messages
	v := NewValidatorWithLang("en")

	// Create a user for benchmark testing
	user := User{
		Username: "john_doe",
		Password: "VerySecurePrd123!",
	}

	// Execute the benchmark test for each language
	for _, lang := range languages {
		// Measure the validation time for each language
		b.Run(lang, func(b *testing.B) {
			// Set the language-specific validation messages

			// Execute the benchmark test
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if err := v.Validate(user); err != nil {
					b.Fatalf("Validation error for language %s: %v", lang, err)
				}
			}
		})
	}
}
