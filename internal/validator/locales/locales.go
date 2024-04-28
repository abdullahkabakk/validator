package locales

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

// ErrorMessages represents a collection of error messages.
type ErrorMessages map[string]string

// LoadMessagesFromJSON loads error messages from a JSON file based on the specified language.
// It reads the JSON file located in the "internal/validator/locales" directory and unmarshal it into an ErrorMessages map.
// The language parameter specifies the language code (e.g., "en" for English).
func LoadMessagesFromJSON(lang string) (ErrorMessages, error) {
	// Get the directory of the package
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("failed to get package directory")
	}
	dir := filepath.Dir(filename)

	// Construct the path to the JSON file based on the specified language
	filename = filepath.Join(dir, lang+".json")

	// Read the content of the JSON file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err // Return the error if unable to read the file
	}

	// Unmarshal the JSON data into the ErrorMessages struct
	var messages ErrorMessages

	if err := json.Unmarshal(data, &messages); err != nil {
		return nil, err // Return the error if unable to unmarshal JSON data
	}

	return messages, nil // Return the loaded error messages
}

// Example usage:
// messages, err := LoadMessagesFromJSON("en")
// if err != nil {
//     fmt.Println("Error loading messages:", err)
// } else {
//     fmt.Println("Messages loaded successfully:", messages)
// }
//
// Available languages:
// - en (English)
// - tr (Turkish)
