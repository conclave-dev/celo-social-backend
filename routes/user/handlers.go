package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stella-zone/celo-social-backend/kvstore"
	"github.com/stella-zone/celo-social-backend/types"
	"github.com/stella-zone/celo-social-backend/util"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	userID := strings.SplitAfter(r.URL.Path, "/user/")[1]

	var user types.User
	var err error
	switch isUser := true; isUser {
	case kvstore.DoesAddressExist(userID):
		user, err = handleExistingAddress(userID)
		if err != nil {
			util.RespondWithError(err, w)
			return
		}

		break
	case kvstore.DoesUserExist(userID):
		user, err = handleExistingName(userID)
		if err != nil {
			util.RespondWithError(err, w)
			return
		}

		break
	case common.IsHexAddress(userID):
		user, err = handleNewAddress(userID)
		if err != nil {
			util.RespondWithError(err, w)
			return
		}

		break
	default:
		util.RespondWithError(errors.New(`Invalid user`), w)
		return
	}

	res, err := json.Marshal(types.JSONResponse{
		Data: user,
	})
	if err != nil {
		util.RespondWithError(err, w)
		return
	}

	w.Write(res)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	userID := strings.SplitAfter(r.URL.Path, "/user/")[1]

	var profile kvstore.Profile
	err := util.ParseJSONBody(r.Body, &profile)
	if err != nil {
		util.RespondWithError(err, w)
		return
	}

	// Validation
	switch isUser := true; isUser {
	// userID path param should always match the profile's "address" or "name" values
	// We should catch this case before reading from cache
	case userID != profile.Address && userID != profile.Name:
		util.RespondWithError(errors.New(`Invalid profile "name" and/or "address"`), w)
		return
	case userID == profile.Address && kvstore.DoesAddressExist(userID):
	case userID == profile.Name && kvstore.DoesUserExist(userID) && kvstore.DoesAddressExist(profile.Address):
		// If either are truthy, we can break out of the switch and create the update
		break
	default:
		util.RespondWithError(errors.New(`Invalid user update"`), w)
		return
	}

	// There should never be a case where the profile "address" is not set (done so at profile creation)
	user, err := fetchAccount(profile.Address)
	if err != nil {
		util.RespondWithError(err, w)
		return
	}

	p, err := json.Marshal(profile)
	if err != nil {
		util.RespondWithError(err, w)
		return
	}

	user.Hash = makeHash(profile, user.AccountSummary, user.Metadata)
	user.Profile = profile

	// Convert marshalled JSON into string and store update
	_, err = kvstore.SetProfile(user.Hash, string(p))
	if err != nil {
		util.RespondWithError(err, w)
		return
	}

	res, err := json.Marshal(types.JSONResponse{
		Data: user,
	})
	if err != nil {
		util.RespondWithError(err, w)
		return
	}

	w.Write(res)
}
