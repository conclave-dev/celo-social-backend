package user

import (
	"github.com/stella-zone/celo-social-backend/kvstore"
	"github.com/stella-zone/celo-social-backend/types"
)

func getUserByAddress(address string) (types.User, error) {
	user, err := fetchAccount(address)
	if err != nil {
		return user, err
	}

	var hash string
	hash = findSocialClaimHash(user.Metadata.Claims)
	if hash == "" || !kvstore.DoesProfileExist(hash) {
		user.Claimed = false

		// If hash is empty, the user may have removed the claim
		// Recover the last claimed hash from their address-hash map
		hash, _ = kvstore.GetAddress(address)
		if hash == "" {

			// Create new profile if no existing address-hash mapping
			// Should only happen the first time a profile is visited
			return createNewProfile(address)
		}
	} else {
		user.Claimed = true
	}

	// Using the hash from the claim or address-hash map, get profile
	profile, err := getProfileFromHash(hash)
	if err != nil {
		return user, err
	}

	// Set to last claimed hash
	_, err = kvstore.SetAddress(address, hash)
	if err != nil {
		return user, err
	}

	user.Claimed = true
	user.Hash = hash
	user.Profile = profile
	return user, nil
}

func getUserByName(username string) (types.User, error) {
	var user types.User
	address, err := kvstore.GetUser(username)
	if err != nil {
		return user, err
	}

	user, err = fetchAccount(address)
	if err != nil {
		return user, err
	}

	hash := findSocialClaimHash(user.Metadata.Claims)
	if hash == "" || !kvstore.DoesProfileExist(hash) {
		user.Claimed = false

		// If hash is empty, the user may have removed the claim
		// Recover the last claimed hash from their address-hash map
		hash, _ = kvstore.GetAddress(address)
		if hash == "" {

			// Create new profile if no existing address-hash mapping
			// Should only happen the first time a profile is visited
			return createNewProfile(address)
		}
	} else {
		user.Claimed = true
	}

	profile, err := getProfileFromHash(hash)
	if err != nil {
		return user, err
	}

	user.Hash = hash
	user.Profile = profile
	return user, nil
}
