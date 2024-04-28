package locales

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

// ErrorMessages represents a collection of error messages.
type ErrorMessages map[string]string

var (
	cache      map[string]ErrorMessages // cache stores loaded error messages for faster retrieval
	cacheMutex sync.RWMutex             // cacheMutex ensures safe concurrent access to the cache
)

func init() {
	cache = make(map[string]ErrorMessages)
}

// LoadMessagesFromJSON loads error messages from a JSON file based on the specified language.
// It reads the JSON file located in the "internal/validator/locales" directory and unmarshal it into an ErrorMessages map.
// The language parameter specifies the language code (e.g., "en" for English).
//
// If the messages for the specified language are already loaded, it returns them from the cache.
// Otherwise, it loads the messages from the JSON file, caches them, and returns.
func LoadMessagesFromJSON(lang string) (ErrorMessages, error) {
	// Check cache first to avoid unnecessary file reads
	cacheMutex.RLock()
	if messages, ok := cache[lang]; ok {
		cacheMutex.RUnlock()
		return messages, nil
	}
	cacheMutex.RUnlock()

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

	// Cache the loaded messages for future use
	cacheMutex.Lock()
	cache[lang] = messages
	cacheMutex.Unlock()

	return messages, nil // Return the loaded error messages
}
