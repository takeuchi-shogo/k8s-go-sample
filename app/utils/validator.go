package utils

import (
	"fmt"
	"net/mail"
)

func ValidateString(value string, min, max int) error {
	n := len(value)
	if n < min || max < n {
		return fmt.Errorf("must contain from %d-%d characters", min, max)
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 12)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not a valid email address")
	}
	return nil
}
