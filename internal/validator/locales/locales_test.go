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

// TestCacheRetrieval tests if messages are retrieved from the cache when available.
func TestCacheRetrieval(t *testing.T) {
	// Initialize cache with some test data
	testLang := "en"
	testMessages := ErrorMessages{
		"key1": "Message 1",
		"key2": "Message 2",
	}
	cache[testLang] = testMessages

	// Perform cache retrieval
	retrievedMessages, err := LoadMessagesFromJSON(testLang)
	if err != nil {
		t.Fatalf("Error while retrieving messages from cache: %v", err)
	}

	// Compare retrieved messages with the expected ones
	if len(retrievedMessages) != len(testMessages) {
		t.Fatalf("Expected %d messages, got %d", len(testMessages), len(retrievedMessages))
	}
	for key, expectedMessage := range testMessages {
		if retrievedMessage, ok := retrievedMessages[key]; !ok || retrievedMessage != expectedMessage {
			t.Errorf("Unexpected message for key %q. Expected: %q, Got: %q", key, expectedMessage, retrievedMessage)
		}
	}
}
