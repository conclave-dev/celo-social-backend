package types

import (
	"net/url"

	"github.com/stella-zone/celo-social-backend/kvstore"
)

func ValidateUser(user kvstore.Profile) {
	if !isURL(user.Photo) {
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
