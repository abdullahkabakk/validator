package validator

import (
	"fmt"
	"github.com/abdullahkabakk/validator/internal/validator/locales"
	"net/mail"
	"reflect"
)

// validateEmail validates if the provided string represents a valid email address format.
// It checks if the email is empty, if it exceeds the maximum length limit (254 characters),
// and if it conforms to the standard email address format.
// If the provided email address is empty, too long, or invalid, it returns an error.
// The fieldName parameter is used to customize the error message to include the name of the field being validated.
// The rule parameter is not currently used in this function but is included for consistency with other validation functions.
func validateEmail(value reflect.Value, messages locales.ErrorMessages, fieldName string, rule string) error {
	// Check if the email is empty
	if value.Len() == 0 {
		return fmt.Errorf(messages["emailIsEmpty"], fieldName)
	}

	// Limit input length to prevent excessive processing time
	// Assuming the maximum length of an email address is 254 characters
	if value.Len() > 254 {
		return fmt.Errorf(messages["emailTooLong"], fieldName)
	}

	// Parse the email address to ensure it conforms to the standard email format
	_, err := mail.ParseAddress(value.String())
	if err != nil {
		return fmt.Errorf(messages["invalidEmail"], fieldName)
	}

	return nil
}
