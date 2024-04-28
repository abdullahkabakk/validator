package validator

import (
	"fmt"
	"github.com/abdullahkabakk/validator/internal/validator/locales"
	"reflect"
	"time"
)

// validateDate validates if the provided string represents a valid date format.
// It checks if the date string conforms to the "YYYY-MM-DD" format and if it represents a valid calendar date.
// If the provided date is not in the correct format or is not a valid date, it returns an error.
// The fieldName parameter is used to customize the error message to include the name of the field being validated.
// The rule parameter is not currently used in this function but is included for consistency with other validation functions.
func validateDate(value reflect.Value, messages locales.ErrorMessages, fieldName string, rule string) error {
	// Limit input length to prevent excessive processing time
	// Assuming the date string is in the format "YYYY-MM-DD", its maximum length is 10 characters.
	if value.Len() > 10 {
		return fmt.Errorf(messages["dateTooLong"], fieldName)
	}

	// Parse the date to ensure it is in the correct format and represents a valid calendar date
	date := value.String()
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return fmt.Errorf(messages["invalidDate"], fieldName)
	}

	return nil
}
