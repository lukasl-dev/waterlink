package internal

import "net/url"

// IsURL returns whether the given string s is a valid URL.
func IsURL(s string) bool {
	_, err := url.Parse(s)
	return err == nil
}
