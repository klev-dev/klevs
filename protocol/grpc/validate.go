package grpc

import (
	"fmt"
)

func (r *LogsListRequest) Validate() error {
	return nil
}

func (r *LogsCreateRequest) Validate() error {
	return ValidateName(r.Name)
}

func (r *LogsDeleteRequest) Validate() error {
	return ValidateName(r.Name)
}

func (r *PublishRequest) Validate() error {
	return ValidateName(r.Name)
}

func (r *ConsumeRequest) Validate() error {
	return ValidateName(r.Name)
}

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
			if i == 0 {
				return fmt.Errorf("at %d: %#U invalid character", i, r)
			}
		default:
			return fmt.Errorf("at %d: %#U invalid character", i, r)
		}
	}

	return nil
}
