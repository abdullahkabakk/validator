package validator

import (
	"reflect"
	"testing"

	"github.com/abdullahkabakk/validator/internal/validator/locales"
)

// TestValidateMinLength tests the validateMinLength function.
func TestValidateMinLength(t *testing.T) {
	tests := []struct {
		value     interface{} // Input value
		minLength string      // Minimum length rule
		expected  bool        // Expected result (true if no error expected, false if error expected)
	}{
		{"hello", "min=7", false},    // value length is less than minLength
		{"world", "min=5", true},     // value length is equal to minLength
		{"greetings", "min=5", true}, // value length is greater than minLength
		// unsupported type
		{[]int{1, 2, 3}, "min=5", false},
		// invalid rule
		{"hello", "invalid", false},
	}

	messages := locales.ErrorMessages{
		"minLength": "Field '%s' must have at least %d characters.",
	}

	for _, test := range tests {
		value := reflect.ValueOf(test.value)
		err := validateMinLength(value, messages, "field", test.minLength)
		// Check if error matches the expectation
		if test.expected && err != nil {
			t.Errorf("Expected no error for value '%v' and minLength %s, got: %v", test.value, test.minLength, err)
		}
		if !test.expected && err == nil {
			t.Errorf("Expected error for value '%v' and minLength %s, got none", test.value, test.minLength)
		}
	}
}

// TestValidateMaxLength tests the validateMaxLength function.
func TestValidateMaxLength(t *testing.T) {
	tests := []struct {
		value     interface{} // Input value
		maxLength string      // Maximum length rule
		expected  bool        // Expected result (true if no error expected, false if error expected)
	}{
		{"hello", "max=5", true},      // value length is less than maxLength
		{"world", "max=5", true},      // value length is equal to maxLength
		{"greetings", "max=5", false}, // value length is greater than maxLength
		// unsupported type
		{[]int{1, 2, 3}, "max=5", false},
		// invalid rule
		{"hello", "invalid", false},
	}

	messages := locales.ErrorMessages{
		"maxLength": "Field '%s' must have at most %d characters.",
	}

	for _, test := range tests {
		value := reflect.ValueOf(test.value)
		err := validateMaxLength(value, messages, "field", test.maxLength)
		// Check if error matches the expectation
		if test.expected && err != nil {
			t.Errorf("Expected no error for value '%v' and maxLength %s, got: %v", test.value, test.maxLength, err)
		}
		if !test.expected && err == nil {
			t.Errorf("Expected error for value '%v' and maxLength %s, got none", test.value, test.maxLength)
		}
	}
}

// TestGetValueLength tests the getValueLength function.
func TestGetValueLength(t *testing.T) {
	tests := []struct {
		name     string      // Test case name
		value    interface{} // Input value
		expected int         // Expected length
		wantErr  bool        // Whether an error is expected
	}{
		{name: "String", value: "hello", expected: 5, wantErr: false},
		{name: "Int", value: 12345, expected: 12345, wantErr: false},
		{name: "Uint", value: uint(54321), expected: 54321, wantErr: false},
		{name: "Float32", value: float32(3.14159), expected: 17, wantErr: false},
		{name: "Float64", value: 3.14159, expected: 7, wantErr: false},
		{name: "Unsupported", value: []int{1, 2, 3}, expected: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := reflect.ValueOf(tt.value)
			got, err := getValueLength(value)
			// Check if error matches the expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("getValueLength() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Check if length matches the expectation
			if got != tt.expected {
				t.Errorf("getValueLength() = %v, want %v", got, tt.expected)
			}
		})
	}
}
