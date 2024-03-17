package api

import (
	"fmt"
	"strings"
)

func ParseLogName(name string) (*LogName, error) {
	parts := strings.Split(name, "/")
	if len(parts) == 1 {
		return NewLogName(parts[0], nil)
	}
	return NewLogName(parts[len(parts)-1], parts[:len(parts)-1])
}

func NewLogName(name string, path []string) (*LogName, error) {
	if err := validate(name); err != nil {
		return nil, fmt.Errorf("invalid name: %w", err)
	}

	for i, p := range path {
		if err := validate(p); err != nil {
			return nil, fmt.Errorf("invalid path element at %d: %w", i, err)
		}
	}

	return &LogName{Name: name, Path: path}, nil
}

func (l *LogName) Validate() error {
	if err := validate(l.Name); err != nil {
		return fmt.Errorf("invalid name: %w", err)
	}

	for i, p := range l.Path {
		if err := validate(p); err != nil {
			return fmt.Errorf("invalid path element at %d: %w", i, err)
		}
	}

	return nil
}

func validate(namePart string) error {
	if len(namePart) == 0 {
		return fmt.Errorf("empty part")
	}

	for i, r := range namePart {
		switch {
		case r >= '0' && r <= '9':
		case r >= 'A' && r <= 'Z':
		case r >= 'a' && r <= 'z':
		case r == '.', r == '_', r == '-':
			if i == 0 {
				return fmt.Errorf("at %d: %#U invalid character", i, r)
			}
		default:
			return fmt.Errorf("at %d: %#U invalid character", i, r)
		}
	}

	return nil
}
