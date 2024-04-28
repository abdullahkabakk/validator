package validator

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/abdullahkabakk/validator/internal/validator/locales"
)

// validateSpecialCharacter validates if a value contains special characters.
// It checks if the value contains any special characters using the containsSpecialCharacter function,
// and returns an error if the value does not contain any special characters.
// The fieldName parameter is used to customize the error message to include the name of the field being validated.
// The rule parameter is not used in this function but is included for consistency with other validation functions.
func validateSpecialCharacter(value reflect.Value, messages locales.ErrorMessages, fieldName string, rule string) error {
	if !containsSpecialCharacter(value.String()) {
		return fmt.Errorf(messages["specialCharacter"], fieldName)
	}
	return nil
}

// containsSpecialCharacter checks if a string contains any special characters.
// It uses a regular expression to match any character that is not alphanumeric.
func containsSpecialCharacter(s string) bool {
	return regexp.MustCompile(`[[:^alnum:]]`).MatchString(s)
}
