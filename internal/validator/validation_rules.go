package validator

import (
	"errors"
	"fmt"
	"github.com/abdullahkabakk/validator/internal/validator/locales"
	"reflect"
	"strconv"
	"strings"
)

// ValidationRule represents a function type for custom validation rules.
// It takes the field value to be validated, error messages, field name, and tag as input, and returns an error if validation fails.
type ValidationRule func(reflect.Value, locales.ErrorMessages, string, string) error

// validationRules maps validation rule names to their corresponding validation functions.
var validationRules = make(map[string]ValidationRule)

// RegisterValidationRule registers a custom validation rule with a given name and validation function.
func registerValidationRule(name string, validateFunc ValidationRule) {
	validationRules[name] = validateFunc
}

// registerDefaultValidationRules registers the default validation rules provided by the package.
func registerDefaultValidationRules() {
	registerValidationRule("required", validateRequired)
	registerValidationRule("min", validateMinLength)
	registerValidationRule("max", validateMaxLength)
	registerValidationRule("uppercase", validateUppercase)
	registerValidationRule("lowercase", validateLowercase)
	registerValidationRule("special", validateSpecialCharacter)
	registerValidationRule("email", validateEmail)
	registerValidationRule("date", validateDate)
}

// validateStruct validates a struct based on the specified validation tags and language.
// It returns an error if validation fails or if any required input is missing.
func validateStruct(input interface{}, lang string) error {
	if input == nil {
		return errors.New("input is nil")
	}

	value := reflect.ValueOf(input)
	if value.Kind() != reflect.Struct {
		return errors.New("input is not a struct")
	}

	// Set the default language to English if not specified
	if lang == "" {
		lang = "en"
	}

	// Load error messages for the specified language
	messages, err := locales.LoadMessagesFromJSON(lang)
	if err != nil {
		// Fallback to English if error messages for the specified language are not available
		messages, err = locales.LoadMessagesFromJSON("en")
		if err != nil {
			return errors.New("failed to load error messages")
		}
	}

	var validationErrors []string

	// Iterate over each struct field and validate based on the validation tags
	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		fieldValue := value.Field(i)
		tag := field.Tag.Get("validate")
		fieldAlias := field.Name

		tags := strings.Split(tag, ",")
		for _, tag := range tags {
			parts := strings.Split(tag, "=")
			// If the tag can be split with '=', it means there is a rule value
			if len(parts) == 2 {
				// Check if the language tag matches the specified language
				if parts[0] == lang {
					fieldAlias = parts[1]
				}
			}
		}

		// Iterate over each tag and apply corresponding validation rules
		for _, tag := range tags {
			// Split the tag into parts
			parts := strings.Split(tag, "=")

			var ruleName string

			// If the tag can be split with '=', it means there is a rule value
			if len(parts) == 2 {
				ruleName = parts[0]
			} else {
				// If it doesn't split, consider the entire tag as the rule name
				ruleName = tag
			}

			// Retrieve the validation function for the rule name
			validateFunc, ok := validationRules[ruleName]
			if !ok {
				// Skip if validation rule is not found
				continue
			}

			// Apply validation function and collect validation errors
			if err := validateFunc(fieldValue, messages, fieldAlias, tag); err != nil {
				validationErrors = append(validationErrors, err.Error())
			}
		}
	}

	// Return concatenated validation errors, if any
	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, ";\n"))
	}

	return nil
}

// parseRule extracts the length from the rule string.
func parseRule(rule string) (int, error) {
	parts := strings.Split(rule, "=")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid rule format: %s", rule)
	}
	return strconv.Atoi(parts[1])
}
