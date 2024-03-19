package klevs

import (
	"fmt"
)

func ValidateName(name string) error {
	if len(name) == 0 {
		return fmt.Errorf("empty part")
	}

	for i, r := range name {
		switch {
		case r >= '0' && r <= '9':
		case r >= 'A' && r <= 'Z':
		case r >= 'a' && r <= 'z':
		case r == '.', r == '_', r == '-':
		default:
			return fmt.Errorf("at %d: %#U invalid character", i, r)
		}
	}

	return nil
}
