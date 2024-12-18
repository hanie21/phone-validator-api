package services

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nyaruka/phonenumbers"
)

// Regular expression to check for ISO 3166-1 alpha-2 format (two uppercase letters)
var isoCodePattern = regexp.MustCompile(`^[A-Z]{2}$`)

// Regular expression to validate spaces between phone number parts
func isValidSpacing(phoneNumber string) bool {
	allowedPattern := `^\+?\d+\s?\d+\s?\d+$`
	re := regexp.MustCompile(allowedPattern)
	return re.MatchString(phoneNumber)
}

// Validate the phone number format and return the parsed number
func ValidatePhoneNumber(phoneNumber, defaultRegion string) (*phonenumbers.PhoneNumber, error) {
	// First, validate that spaces are in the right places
	if !isValidSpacing(phoneNumber) {
		return nil, fmt.Errorf("invalid space or character usage in phone number")
	}

	// Remove all spaces before passing to libphonenumber for parsing
	phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")

	// If the phone number doesn't start with +, prepend +, as the phonenumbers package need + to recognize country code
	if !strings.HasPrefix(phoneNumber, "+") && defaultRegion == "" {
		phoneNumber = "+" + phoneNumber
	}

	// First, attempt to parse the phone number without assuming it's missing a country code
	num, err := phonenumbers.Parse(phoneNumber, defaultRegion)
	if err != nil {
		// If parsing failed, check if a valid `countryCode` is provided
		if defaultRegion != "" {
			// Validate if the countryCode is in ISO 3166-1 alpha-2 format
			if !isoCodePattern.MatchString(defaultRegion) {
				return nil, fmt.Errorf("invalid country code format, must be ISO 3166-1 alpha-2")
			}

			// Retry parsing with the provided ISO 3166-1 country code, also remove + from phone number
			num, err = phonenumbers.Parse(phoneNumber[1:], defaultRegion)
			if err != nil {
				return nil, fmt.Errorf("error parsing phone number with country code")
			}
		} else {
			return nil, fmt.Errorf("error parsing phone number and no country code provided")
		}
	}

	// Validate that the parsed number is correct
	if !phonenumbers.IsValidNumber(num) {
		return nil, fmt.Errorf("invalid phone number, required value is missing")
	}

	return num, nil
}

// Extract phone number components like country code, area code, and local number
func ExtractPhoneNumberComponents(num *phonenumbers.PhoneNumber) (string, string, string, string) {
	// Format the number in E.164 format
	e164Format := phonenumbers.Format(num, phonenumbers.E164)

	// Get the country code and national number
	countryCodeStr := phonenumbers.GetRegionCodeForNumber(num)
	nationalNumber := phonenumbers.GetNationalSignificantNumber(num)
	areaCode := nationalNumber[:3]         // First 3 digits as area code
	localPhoneNumber := nationalNumber[3:] // Remaining digits as local phone number

	return e164Format, countryCodeStr, areaCode, localPhoneNumber
}
