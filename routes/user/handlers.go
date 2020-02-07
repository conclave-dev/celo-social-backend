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
	"github.com/stella-zone/celo-social-backend/types"
	"github.com/stella-zone/celo-social-backend/util"
)

const apiURL = "http://localhost:8080"

func getUser(w http.ResponseWriter, r *http.Request) {
	userID := strings.SplitAfter(r.URL.Path, "/user/")[1]

	switch isUser := true; isUser {
	case kvstore.DoesAddressExist(userID):
		fmt.Print("address exists")
		break
	case kvstore.DoesUserExist(userID):
		fmt.Print("user exists")
		break
	case common.IsHexAddress(userID):
		fmt.Print("is hex address")
		// Fetch account summary
		// Check metadata for claims
		// If claim(s) exist and formatted, create an Address and User
		// Else, respond with account summary
		break
	default:
		fmt.Print("no matches")
		return
	}

	w.Write([]byte("hi"))
	return
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var p types.User
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
		Data: types.UpdateResponse{
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
