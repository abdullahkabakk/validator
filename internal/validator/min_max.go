package validator

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/abdullahkabakk/validator/internal/validator/locales"
)

// validateMinLength validates if the length of a value is greater than or equal to the minimum length.
// It parses the rule string to extract the minimum length requirement, then compares it with the length of the value.
// If the length of the value is less than the minimum length, it returns an error.
// The fieldName parameter is used to customize the error message to include the name of the field being validated.
// The rule parameter specifies the minimum length requirement.
func validateMinLength(value reflect.Value, messages locales.ErrorMessages, fieldName string, rule string) error {
	// Parse the rule to get the minimum length
	minLength, err := parseRule(rule)
	if err != nil {
		return err
	}

	// Get the length of the value
	length, err := getValueLength(value)
	if err != nil {
		return err
	}

	// Check if the length is less than the minimum length
	if length < minLength {
		return fmt.Errorf(messages["minLength"], fieldName, minLength)
	}

	return nil
}

// validateMaxLength validates if the length of a value is less than or equal to the maximum length.
// It parses the rule string to extract the maximum length requirement, then compares it with the length of the value.
// If the length of the value exceeds the maximum length, it returns an error.
// The fieldName parameter is used to customize the error message to include the name of the field being validated.
// The rule parameter specifies the maximum length requirement.
func validateMaxLength(value reflect.Value, messages locales.ErrorMessages, fieldName string, rule string) error {
	// Parse the rule to get the maximum length
	maxLength, err := parseRule(rule)
	if err != nil {
		return err
	}

	// Get the length of the value
	length, err := getValueLength(value)
	if err != nil {
		return err
	}

	// Check if the length is greater than the maximum length
	if length > maxLength {
		return fmt.Errorf(messages["maxLength"], fieldName, maxLength)
	}

	return nil
}

// getValueLength returns the length of a value based on its type.
// It supports string, integer, floating-point, and unsigned integer types.
// For other types, it returns an error indicating that the type is not supported for length validation.
func getValueLength(value reflect.Value) (int, error) {
	// Check the type of the value and get its length
	switch value.Kind() {
	case reflect.String:
		return len(value.String()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(value.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int(value.Uint()), nil
	case reflect.Float32, reflect.Float64:
		str := strconv.FormatFloat(value.Float(), 'f', -1, 64)
		return len(str), nil
	default:
		return 0, fmt.Errorf("unsupported type for length validation: %v", value.Kind())
	}
}
