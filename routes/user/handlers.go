package user

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stella-zone/celo-social-backend/kvstore"
	_types "github.com/stella-zone/celo-social-backend/types"
	"github.com/stella-zone/celo-social-backend/util"
	"github.com/stella-zone/go-celo/types"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	userID := strings.SplitAfter(r.URL.Path, "/user/")[1]

	var user _types.User
	switch isUser := true; isUser {
	case kvstore.DoesAddressExist(userID):
		fmt.Print("address exists")
		break
	case kvstore.DoesUserExist(userID):
		fmt.Print("user exists")
		break
	case common.IsHexAddress(userID):
		handleUnclaimedUser(userID, &user, w)

		break
	default:
		util.RespondWithError(errors.New(`User not found with the specified "address" or "username"`), w)
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
	return
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var p _types.Profile
	err := util.ParseJSONBody(w, r.Body, &p)
	if err != nil {
		util.RespondWithError(err, w)
		return
	}

	// Convert decoded request body to sha256 checksum
	checkSum := sha256.Sum256([]byte(fmt.Sprint(p)))
	hash := fmt.Sprint(hex.EncodeToString(checkSum[:]))

	// Check if the update already exists
	exists := kvstore.DoesUpdateExist(hash)
	if exists == true {
		err := errors.New("Update already exists")
		util.RespondWithError(err, w)
		return
	}

	update, err := json.Marshal(p)
	if err != nil {
		util.RespondWithError(err, w)
		return
	}

	// Convert marshalled JSON into string and store update
	kvstore.SetUpdate(hash, string(update))

	res, err := json.Marshal(types.JSONResponse{
		Data: _types.UpdateUserResponse{
			Hash:   hash,
			Update: string(update),
		},
	})
	if err != nil {
		util.RespondWithError(err, w)
		return
	}

	w.Write(res)
}
