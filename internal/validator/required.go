package validator

import (
	"fmt"
	"reflect"

	"github.com/abdullahkabakk/validator/internal/validator/locales"
)

// validateRequired validates if a value is not empty.
// It checks if the value is empty using the isEmpty function,
// and returns an error if the value is empty.
// The fieldName parameter is used to customize the error message to include the name of the field being validated.
// The rule parameter is not used in this function but is included for consistency with other validation functions.
func validateRequired(value reflect.Value, messages locales.ErrorMessages, fieldName string, rule string) error {
	if isEmpty(value) {
		return fmt.Errorf(messages["required"], fieldName)
	}
	return nil
}

// isEmpty checks if a value is empty.
// It supports string, slice, and map types.
// For non-string types, it checks if the value is equal to its zero value.
func isEmpty(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Slice, reflect.Map:
		return value.Len() == 0
	default:
		// For non-string types, check if the value is equal to its zero value.
		return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
	}
}
