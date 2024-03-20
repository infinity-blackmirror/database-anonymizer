package faker

import (
	"testing"
)

func TestIsValidFaker(t *testing.T) {
	manager := NewFakeManager()

	if !manager.IsValidFaker("") {
		t.Fatalf("TestIsValidFaker: empty faker check failed")
	}

	if !manager.IsValidFaker("_") {
		t.Fatalf("TestIsValidFaker: _ faker check failed")
	}

	if !manager.IsValidFaker("address") {
		t.Fatalf("TestIsValidFaker: address faker check failed")
	}

	if manager.IsValidFaker("unknown_faker") {
		t.Fatalf("TestIsValidFaker: unknown_faker faker check failed")
	}
}
