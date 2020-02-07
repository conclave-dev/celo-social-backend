package types

import (
	"net/url"
)

func ValidateUser(user Profile) {
	if !isURL(user.PhotoURL) {
		// Respond with error
	}
}

func isURL(rawurl string) bool {
	_, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return false
	}

	return true
}
