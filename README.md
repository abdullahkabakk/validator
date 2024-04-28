# Struct Validator: Effortless Struct Field Validation in Go

Struct Validator is a versatile Go package designed to streamline the validation of struct fields using intuitive validation tags. Whether you're building a web application, API, or any other Go-based project, Struct Validator simplifies the task of ensuring your data meets specific criteria.

Struct Validator allows you to define validation rules using struct tags, making it easy to validate struct fields with minimal effort. You can define custom validation rules, support multiple languages for error messages, and validate structs with default or custom languages.

Struct Validator is designed to be flexible, customizable, and easy to use, making it the perfect choice for any Go project that requires robust data validation.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Custom Validation Rules](#custom-validation-rules)
- [Multilingual Support](#multilingual-support)
- [License](#license)

## Features

- **Validation Tags:** Define validation rules using struct tags.
- **Customizable:** Easily define custom validation rules.
- **Multilingual Support:** Supports validation error messages in multiple languages.
- **Flexible:** Validate structs with default or custom languages.

## Installation

Install Struct Validator using `go get`:

```bash
go get github.com/abdullahkabakk/validator
```

To use Struct Validator in your Go project, simply import it:

```bash
import "github.com/abdullahkabakk/validator"
```

## Usage

To use Struct Validator, you need to define validation rules using struct tags. Here's an example:

```go
package main

import (
    "fmt"
    "github.com/abdullahkabakk/validator"
)

type User struct {
    Name     string `validate:"required"`
    Email    string `validate:"required,email"`
    Password string `validate:"required,min=8"`
}

func main() {
    user := User{
        Name:     "John Doe",
        Email:    "john@doe.com",
        Password: "password",
    }

    v := validator.NewValidator()
    if err := v.Validate(user); err != nil {
        fmt.Println(err)
    }
}
```

### Explanation:

- **Creating User Struct**: We define a `User` struct with three fields: `Name`, `Email`, and `Password`. Each field has a validation tag that specifies the validation rules.

- **Initializing Validator**: We create a new instance of the `Validator` struct using `validator.NewValidator()`.

- **Validating User Struct**: We validate the `User` struct using the `Validate` method of the `Validator` instance. If any validation rule fails, an error message will be returned.

In this example, we define a `User` struct with three fields: `Name`, `Email`, and `Password`. We then define validation rules for each field using struct tags. The `Name` field is required, the `Email` field must be a valid email address, and the `Password` field must be at least 8 characters long.

We then create a new `Validator` instance and call the `Validate` method with the `User` struct. If any of the validation rules fail, an error message will be returned.

## Custom Validation Rules

You can define custom validation rules by implementing the `ValidatorFunc` interface. Here's an example:

```go
package main

import (
	"fmt"
	"github.com/abdullahkabakk/validator"
)

type User struct {
	Name string `validate:"required"`
	Age  int    `validate:"required,custom=validateAge"`
}

func validateAge(value interface{}) error {
	age := value.(int)
	if age < 18 {
		return fmt.Errorf("age must be at least 18")
	}
	return nil
}

func main() {
	// Define a User instance
	user := User{
		Name: "John Doe",
		Age:  16,
	}

	// Create a new Validator instance
	v := validator.NewValidator()

	// Register the custom validation rule
	v.RegisterValidator("validateAge", validateAge)

	// Validate the User struct
	if err := v.Validate(user); err != nil {
		fmt.Println(err)
	}
}
```

### Explanation:

- **Creating User Struct**: We define a `User` struct with two fields: `Name` and `Age`. The `Name` field is required, and the `Age` field has a custom validation rule called `validateAge`.

- **Custom Validation Rule**: We define a custom validation rule called `validateAge` that checks if the `Age` field is at least 18. The custom validation rule takes an `interface{}` value as input and returns an error if the validation fails.

- **Initializing Validator**: We create a new instance of the `Validator` struct using `validator.NewValidator()`.

- **Registering Custom Validation Rule**: We register the custom validation rule `validateAge` with the `Validator` instance using the `RegisterValidator` method.

- **Validating User Struct**: We validate the `User` struct using the `Validate` method of the `Validator` instance. If the custom validation rule fails, an error message will be returned.

In this example, we define a custom validation rule called `validateAge` that checks if the `Age` field is at least 18. We then register this custom validation rule with the `Validator` instance using the `RegisterValidator` method.

## Multilingual Support

Struct Validator supports validation error messages in multiple languages. You can set the language for the error messages using the `SetLanguage` method. Here's an example:

```go

package main

import (
    "fmt"
    "github.com/abdullahkabakk/validator"
)

type User struct {
    Name     string `validate:"required"`
    Email    string `validate:"required,email"`
    Password string `validate:"required,min=8"`
}

func main() {
    user := User{
        Name:     "John Doe",
        Email:    "john@doe.com",
        Password: "password",
    }

    v := validator.NewValidator()
    v.SetLang("tr")

    if err := v.Validate(user); err != nil {
        fmt.Println(err)
    }

    v.SetLang("en")

    if err := v.Validate(user); err != nil {
        fmt.Println(err)
    }

}
```

### Explanation:

- **Creating User Instance**: First, we define a `User` instance with sample data for demonstration purposes.

- **Initializing Validator**: We create a new instance of the `Validator` struct using `validator.NewValidator()`.

- **Setting Language to Turkish**: We set the language for the error messages to Turkish using the `SetLang` method with the language code `"tr"`.

- **Validating User Struct**: We then validate the `User` struct using the `Validate` method of the `Validator` instance. If any validation rule fails, the error messages will be returned in Turkish.

- **Switching Language to English**: Next, we switch the language to English by calling `SetLang("en")`.

- **Validating User Struct Again**: We validate the `User` struct once more. If any validation rule fails this time, the error messages will be returned in English.

In this example, we demonstrate how to set the language for error messages in Struct Validator. You can switch between languages using the `SetLang` method, which accepts a language code as input.

## License

Struct Validator is licensed under the MIT license. See the [LICENSE](LICENSE) file for more information.
