package pkgerror

import "fmt"

// New returns a new package error.
func New(msg string) error {
	return fmt.Errorf("waterlink: %s", msg)
}

// Wrap wraps err in a new package error.
func Wrap(prefix string, err error) error {
	if err == nil {
		return nil
	}
	return New(fmt.Sprintf("%s: %s", prefix, err.Error()))
}
