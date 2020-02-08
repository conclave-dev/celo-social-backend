package user

import (
	"encoding/json"
	"net/http"

	"github.com/stella-zone/celo-social-backend/kvstore"
	"github.com/stella-zone/celo-social-backend/types"
	"github.com/stella-zone/celo-social-backend/util"
)

func handleExistingAddress(address string, w http.ResponseWriter) (types.User, error) {
	user, err := fetchAccount(address, w)
	if err != nil {
		util.RespondWithError(err, w)
		return user, err
	}

	hash, err := kvstore.GetAddress(address)
	if err != nil {
		util.RespondWithError(err, w)
		return user, err
	}

	// TODO: How do we handle existing claims?
	// What if the address hash doesn't map to a profile?
	// How do we ensure that the user and address hashes match?

	profile, err := kvstore.GetProfile(hash)
	if err != nil {
		util.RespondWithError(err, w)
		return user, err
	}

	if profile == "" {
		// TODO: How do we handle an address hash that doesn't map to a profile?
		// What would cause an address hash to not map to a profile?
	}

	var p kvstore.Profile
	err = json.Unmarshal([]byte(profile), &p)
	if err != nil {
		util.RespondWithError(err, w)
		return user, err
	}

	socialClaims := findSocialClaims(user.Metadata.Claims)
	if len(socialClaims) == 0 {
		var err error
		_, err = kvstore.DeleteUser(p.Name)
		_, err = kvstore.DeleteAddress(address)
		_, err = kvstore.DeleteProfile(hash)
		return user, err
	}

	user.Hash = hash
	user.Profile = p

	return user, nil
}

func handleExistingName(username string, w http.ResponseWriter) (types.User, error) {
	var user types.User
	hash, err := kvstore.GetUser(username)
	if err != nil {
		util.RespondWithError(err, w)
		return user, err
	}

	profile, err := kvstore.GetProfile(hash)
	if err != nil {
		util.RespondWithError(err, w)
		return user, err
	}

	if profile == "" {
		// TODO: How do we handle an address hash that doesn't map to a profile?
		// What would cause an address hash to not map to a profile?
	}

	var p kvstore.Profile
	err = json.Unmarshal([]byte(profile), &p)
	if err != nil {
		util.RespondWithError(err, w)
		return user, err
	}

	user, err = fetchAccount(p.Address, w)
	if err != nil {
		util.RespondWithError(err, w)
		return user, err
	}

	socialClaims := findSocialClaims(user.Metadata.Claims)
	if len(socialClaims) == 0 {
		_, err = kvstore.DeleteUser(username)
		_, err = kvstore.DeleteAddress(p.Address)
		_, err = kvstore.DeleteProfile(hash)
		return user, err
	}

	user.Hash = hash
	user.Profile = p

	// TODO: How do we handle existing claims?
	// What if the user hash doesn't map to a profile?
	// How do we ensure that the user and address hashes match?

	return user, nil
}

func handleNewAddress(address string, w http.ResponseWriter) (types.User, error) {
	user, err := fetchAccount(address, w)
	if err != nil {
		util.RespondWithError(err, w)
		return user, err
	}

	socialClaims := findSocialClaims(user.Metadata.Claims)
	if len(socialClaims) == 0 {
		return user, nil
	}

	// Create new user profile and marshal into storable string type
	user.Profile = MakeProfile(user)
	p, err := json.Marshal(user.Profile)
	if err != nil {
		util.RespondWithError(err, w)
		return user, err
	}

	// Generate a hash from user profile and account data, and set as key to profile
	user.Hash = MakeHash(user.Profile, user.AccountSummary, user.Metadata)
	_, err = kvstore.SetProfile(user.Hash, string(p))
	if err != nil {
		util.RespondWithError(err, w)
		return user, err
	}

	// Map address to hash, to get profile later (i.e. address -> hash -> profile)
	_, err = kvstore.SetAddress(user.AccountSummary.Address.String(), user.Hash)
	if err != nil {
		util.RespondWithError(err, w)
		return user, err
	}

	// Celo accounts are not required to have names, so we stop here if not set
	if user.AccountSummary.Name == "" {
		return user, nil
	}

	// Map name to hash, to get profile later (i.e. name -> hash -> profile)
	_, err = kvstore.SetUser(user.AccountSummary.Name, user.Hash)
	if err != nil {
		util.RespondWithError(err, w)
		return user, err
	}

	return user, nil
}
