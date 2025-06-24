package main

import "testing"

func TestThemeValidate(t *testing.T) {
	testTheme := newRandomTheme(false)

	if err := testTheme.Validate(); err == nil {
		t.Error("theme passed validation without required fields")
	}

	testTheme.Name = "My Theme"

	if err := testTheme.Validate(); err != nil {
		t.Error("theme failed validation")
	}
}
