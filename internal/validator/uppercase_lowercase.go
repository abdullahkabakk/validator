package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/abdullahkabakk/validator/internal/validator/locales"
)

// validateUppercase validates if a value contains at least one uppercase letter.
// It checks if the value contains at least one uppercase letter using the containsUppercase function,
// and returns an error if the value does not contain any uppercase letters.
// The fieldName parameter is used to customize the error message to include the name of the field being validated.
// The rule parameter is not used in this function but is included for consistency with other validation functions.
func validateUppercase(value reflect.Value, messages locales.ErrorMessages, fieldName string, rule string) error {
	if !containsUppercase(value.String()) {
		return fmt.Errorf(messages["uppercaseLetter"], fieldName)
	}
	return nil
}

// validateLowercase validates if a value contains at least one lowercase letter.
// It checks if the value contains at least one lowercase letter using the containsLowercase function,
// and returns an error if the value does not contain any lowercase letters.
// The fieldName parameter is used to customize the error message to include the name of the field being validated.
// The rule parameter is not used in this function but is included for consistency with other validation functions.
func validateLowercase(value reflect.Value, messages locales.ErrorMessages, fieldName string, rule string) error {
	if !containsLowercase(value.String()) {
		return fmt.Errorf(messages["lowercaseLetter"], fieldName)
	}
	return nil
}

// containsUppercase checks if a string contains at least one uppercase letter.
// It compares the string with its lowercase representation to determine if any characters were uppercase.
func containsUppercase(s string) bool {
	return strings.ToLower(s) != s
}

// containsLowercase checks if a string contains at least one lowercase letter.
// It compares the string with its uppercase representation to determine if any characters were lowercase.
func containsLowercase(s string) bool {
	return strings.ToUpper(s) != s
}
