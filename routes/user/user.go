package user

import (
	"encoding/json"
	"net/http"

	"github.com/stella-zone/celo-social-backend/kvstore"
	"github.com/stella-zone/celo-social-backend/types"
	"github.com/stella-zone/celo-social-backend/util"
)

func handleNewAddress(address string, w http.ResponseWriter) (types.User, error) {
	user, err := fetchAccount(address, w)
	if err != nil {
		util.RespondWithError(err, w)
		return user, err
	}

	if len(user.Metadata.Claims) == 0 {
		return user, nil
	}

	// Create new user profile and marshal into storable string type
	user.Profile = makeProfile()
	p, err := json.Marshal(user.Profile)
	if err != nil {
		util.RespondWithError(err, w)
		return user, err
	}

	// Generate a hash from user profile and account data, and set as key to profile
	user.Hash = makeHash(user.Profile, user.AccountSummary, user.Metadata)
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
