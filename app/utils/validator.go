package utils

import "fmt"

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
