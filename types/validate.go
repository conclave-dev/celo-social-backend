package types

import (
	"net/url"
)

func ValidateProfile(profile Profile) {
	if !isURL(profile.PhotoURL) {
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
