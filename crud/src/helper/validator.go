package helper

import (
	"errors"
	"fmt"
	"regexp"
)

// validateInput validates the input fields based on specified criteria.
func ValidateInput(firstname, lastname, email, phone string, id *int) error {
	// Validate firstname (1-30 characters)
	if !IsValidName(firstname) {
		return ErrInvalidFirstname
	}

	// Validate lastname (1-30 characters)
	if !IsValidName(lastname) {
		return ErrInvalidLastname
	}

	// Validate email (valid email format)
	if !IsValidEmail(email) {
		return ErrInvalidEmail
	}

	// Validate phone (valid phone number format)
	if !IsValidPhone(phone) {
		return ErrInvalidPhone
	}

	return nil
}
func ValidateUpdateUserInput(fields map[string]interface{}) error {

	if val, ok := fields["id"]; ok {
		if val.(int) < 0 {
			return ErrInvalidID
		}
	}
	if val, ok := fields["firstname"]; ok {
		if !IsValidName(fmt.Sprintf("%v", val)) {
			return ErrInvalidFirstname
		}
	}

	if val, ok := fields["lasttname"]; ok {
		if !IsValidName(fmt.Sprintf("%v", val)) {
			return ErrInvalidLastname
		}
	}
	if val, ok := fields["email"]; ok {
		if !IsValidEmail(fmt.Sprintf("%v", val)) {
			return ErrInvalidEmail
		}
	}

	if val, ok := fields["phone"]; ok {
		if !IsValidPhone(fmt.Sprintf("%v", val)) {
			return ErrInvalidPhone
		}
	}
	return nil
}

// isValidEmail checks if the provided email is in a valid format.
func IsValidEmail(email string) bool {
	// Regular expression for basic email validation
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailRegex, email)
	return matched
}
func IsValidName(name string) bool {
	// Regular expression for basic email validation
	nameRgx := `^[a-zA-Z ]{1,30}$`
	matched, _ := regexp.MatchString(nameRgx, name)
	return matched
}

// isValidPhone checks if the provided phone number is in a valid format.
func IsValidPhone(phone string) bool {
	// Regular expression for basic phone number validation (allowing only digits and dashes)
	phoneRegex := `^[0-9-]+$`
	matched, _ := regexp.MatchString(phoneRegex, phone)
	return matched
}

// Custom error messages for validation
var (
	ErrInvalidFirstname = errors.New("firstname must be 1-30 characters long")
	ErrInvalidLastname  = errors.New("lastname must be 1-30 characters long")
	ErrInvalidEmail     = errors.New("invalid email format")
	ErrInvalidPhone     = errors.New("invalid phone number format")
	ErrInvalidID        = errors.New("invalid ID")
	ErrUserNotFound     = errors.New("User doesn't exists")
)
