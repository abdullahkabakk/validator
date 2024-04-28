package validator

import (
	"errors"
	"github.com/abdullahkabakk/validator/internal/validator/locales"
	"reflect"
	"testing"
)

// TestNewValidator tests the NewValidator function.
func TestNewValidator(t *testing.T) {
	// Create a new validator instance
	v := NewValidator()

	// Check if the default language is set to English
	if v.DefaultLang != "en" {
		t.Errorf("Expected default language to be 'en', got '%s'", v.DefaultLang)
	}
}

// TestNewValidatorWithLang tests the NewValidatorWithLang function.
func TestNewValidatorWithLang(t *testing.T) {
	// Create a new validator instance with a specified default language
	defaultLang := "tr"
	v := NewValidatorWithLang(defaultLang)

	// Check if the default language is set correctly
	if v.DefaultLang != defaultLang {
		t.Errorf("Expected default language to be '%s', got '%s'", defaultLang, v.DefaultLang)
	}
}

// TestValidate tests the Validate function.
func TestValidate(t *testing.T) {
	// Define a struct for testing
	type User struct {
		Username string `validate:"required,min=3,max=20"`
		Email    string `validate:"required,email"`
		Password string `validate:"required,min=8"`
	}

	// Create a validator instance
	validator := NewValidator()

	// Valid user
	user := User{Username: "john_doe", Email: "john@example.com", Password: "12345678"}
	err := validator.Validate(user)
	if err != nil {
		t.Errorf("Expected validation to pass, got error: %v", err)
	}

	// Invalid user (missing required fields)
	invalidUser := User{}
	err = validator.Validate(invalidUser)
	if err == nil {
		t.Errorf("Expected validation to fail, but it passed")
	}
}

// TestValidateWithLang tests the ValidateWithLang function.
func TestValidateWithLang(t *testing.T) {
	// Define a struct for testing
	type User struct {
		Username string `validate:"required,min=3,max=20"`
		Email    string `validate:"required,email"`
		Password string `validate:"required,min=8"`
	}

	// Create a validator instance
	validator := NewValidator()

	// Valid user
	user := User{Username: "john_doe", Email: "john@example.com", Password: "12345678"}
	err := validator.ValidateWithLang(user, "en")
	if err != nil {
		t.Errorf("Expected validation to pass, got error: %v", err)
	}

	// Invalid user (missing required fields)
	invalidUser := User{}
	err = validator.ValidateWithLang(invalidUser, "en")
	if err == nil {
		t.Errorf("Expected validation to fail, but it passed")
	}
}

// TestSetLang tests the SetLang function.
func TestSetLang(t *testing.T) {
	// Create a new validator instance
	v := NewValidator()

	// Set the default language to Turkish
	v.SetLang("tr")

	// Check if the default language is set correctly
	if v.DefaultLang != "tr" {
		t.Errorf("Expected default language to be 'tr', got '%s'", v.DefaultLang)
	}
}

// TestRegisterValidationRule tests the RegisterValidationRule function.
func TestRegisterValidationRule(t *testing.T) {
	// Create a new validator instance
	v := NewValidator()

	// Define a custom validation rule
	validateCustom := func(value reflect.Value, messages locales.ErrorMessages, fieldValue string, rule string) error {
		if value.Int() < 0 {
			return errors.New("Value must be greater than or equal to 0")
		}
		return nil
	}

	// Register the custom validation rule
	v.RegisterValidationRule("custom", validateCustom)

	// Define a struct with a custom validation rule
	type Data struct {
		Value int `validate:"custom"`
	}

	// Valid data
	validData := Data{Value: 10}
	err := v.Validate(validData)
	if err != nil {
		t.Errorf("Expected validation to pass, got error: %v", err)
	}
}
