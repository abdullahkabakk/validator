package validator

import "github.com/abdullahkabakk/validator/internal/validator"

// Validator represents a validation instance that can be used to validate structs.
type Validator struct {
	// DefaultLang holds the default language for validation error messages.
	DefaultLang string
}

// NewValidator creates a new instance of Validator with the default language set to English.
func NewValidator() *Validator {
	return NewValidatorWithLang("en")
}

// NewValidatorWithLang creates a new instance of Validator with the specified default language.
func NewValidatorWithLang(defaultLang string) *Validator {
	return &Validator{
		DefaultLang: defaultLang,
	}
}

// Validate performs validation on the input struct using the default language.
// It registers default validation rules and then validates the struct fields.
func (v *Validator) Validate(input interface{}) error {
	return v.ValidateWithLang(input, v.DefaultLang)
}

// ValidateWithLang performs validation on the input struct using the specified language.
// It validates the struct fields based on the validation tags and returns any validation errors encountered.
func (v *Validator) ValidateWithLang(input interface{}, lang string) error {
	validator.RegisterDefaultValidationRules()
	return validator.ValidateStruct(input, lang)
}

// SetLang sets the default language for validation error messages.
func (v *Validator) SetLang(lang string) {
	v.DefaultLang = lang
}

// RegisterValidationRule registers a custom validation rule with a given name and validation function.
func (v *Validator) RegisterValidationRule(name string, validateFunc validator.ValidationRule) {
	validator.RegisterValidationRule(name, validateFunc)
}

// Example usage:
//
//   type User struct {
//       Username string `validate:"required,min=3,max=20"`
//       Email    string `validate:"required,email"`
//       Password string `validate:"required,min=8"`
//   }
//
//   func main() {
//       validator := NewValidator()
//       user := User{Username: "john_doe", Email: "john@example.com", Password: "12345678"}
//       if err := validator.Validate(user); err != nil {
//           fmt.Println("Validation error:", err)
//       } else {
//           fmt.Println("Validation passed")
//       }
//   }
