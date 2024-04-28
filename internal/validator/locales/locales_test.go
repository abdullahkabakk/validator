package locales

import (
	"testing"
)

// TestLoadMessagesFromJSON tests the LoadMessagesFromJSON function.
func TestLoadMessagesFromJSON(t *testing.T) {
	testCases := []struct {
		lang        string
		expectError bool
	}{
		{lang: "en", expectError: false},          // Valid language
		{lang: "invalid_lang", expectError: true}, // Invalid language
	}

	for _, tc := range testCases {
		t.Run(tc.lang, func(t *testing.T) {
			// Execute the test case
			messages, err := LoadMessagesFromJSON(tc.lang)

			// Check if an error is expected
			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error for language %s, but got none", tc.lang)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for language %s: %v", tc.lang, err)
				}
				if len(messages) == 0 {
					t.Errorf("No messages loaded for language %s", tc.lang)
				}
			}
		})
	}
}
