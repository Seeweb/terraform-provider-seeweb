package seeweb

import (
	"testing"
)

// Test config with an empty token
func TestConfigEmptyToken(t *testing.T) {
	config := Config{
		Token: "",
	}

	if _, err := config.Client(); err == nil {
		t.Fatalf("expected error, but got nil")
	}
}

// Test config with a custom ApiUrl override
func TestConfigCustomApiUrlOverride(t *testing.T) {
	config := Config{
		Token:          "foo",
		ApiUrlOverride: "https://api.domain-override.tld",
	}

	if _, err := config.Client(); err != nil {
		t.Fatalf("error: expected the client to not fail: %v", err)
	}
}
