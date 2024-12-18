package services

import (
	"testing"

	"github.com/nyaruka/phonenumbers"
)

// Test case for valid phone numbers with and without country code
func TestValidatePhoneNumber(t *testing.T) {
	// Test cases
	tests := []struct {
		phoneNumber   string
		defaultRegion string
		expectedError bool
	}{
		{phoneNumber: "+34915872200", defaultRegion: "", expectedError: false},    // Valid Spanish number with country code
		{phoneNumber: "915872200", defaultRegion: "ES", expectedError: false},     // Valid Spanish number without country code
		{phoneNumber: "+12125690123", defaultRegion: "", expectedError: false},    // Valid US number with country code
		{phoneNumber: "2125690123", defaultRegion: "US", expectedError: false},    // Valid US number without country code
		{phoneNumber: "+34345 872200", defaultRegion: "", expectedError: true},    // Invalid phone number with wrong space
		{phoneNumber: "invalidNumber", defaultRegion: "", expectedError: true},    // Invalid non-numeric phone number
		{phoneNumber: "915872200", defaultRegion: "", expectedError: true},        // Valid local number but no defaultRegion
		{phoneNumber: "915872200", defaultRegion: "INVALID", expectedError: true}, // Invalid country code
	}

	// Loop through test cases
	for _, test := range tests {
		num, err := ValidatePhoneNumber(test.phoneNumber, test.defaultRegion)

		if test.expectedError && err == nil {
			t.Errorf("Expected error for phoneNumber: %s, but got none", test.phoneNumber)
		} else if !test.expectedError && err != nil {
			t.Errorf("Unexpected error for phoneNumber: %s - %v", test.phoneNumber, err)
		} else if !test.expectedError {
			if !phonenumbers.IsValidNumber(num) {
				t.Errorf("Expected valid number for phoneNumber: %s, but it was invalid", test.phoneNumber)
			}
		}
	}
}

// Test case for extracting phone number components
func TestExtractPhoneNumberComponents(t *testing.T) {
	// Parse a valid phone number
	num, err := ValidatePhoneNumber("+34915872200", "")
	if err != nil {
		t.Fatalf("Failed to validate phone number: %v", err)
	}

	e164Format, countryCodeStr, areaCode, localPhoneNumber := ExtractPhoneNumberComponents(num)

	// Check if the extracted values match expectations
	if e164Format != "+34915872200" {
		t.Errorf("Expected E.164 format: +34915872200, but got: %s", e164Format)
	}
	if countryCodeStr != "ES" {
		t.Errorf("Expected country code: ES, but got: %s", countryCodeStr)
	}
	if areaCode != "915" {
		t.Errorf("Expected area code: 915, but got: %s", areaCode)
	}
	if localPhoneNumber != "872200" {
		t.Errorf("Expected local phone number: 872200, but got: %s", localPhoneNumber)
	}
}
