package profile

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/stella-zone/celo-social-backend/kvstore"
	"github.com/stella-zone/celo-social-backend/types"
	"github.com/stella-zone/celo-social-backend/util"
)

func updateProfile(w http.ResponseWriter, r *http.Request) {
	var p types.Profile
	err := util.ParseRequestParameters(w, r, &p)
	if err != nil {
		util.RespondWithError(err, w)
		return
	}

	// Convert decoded request body to sha256 checksum
	checkSum := sha256.Sum256([]byte(fmt.Sprint(p)))
	hash := fmt.Sprint(hex.EncodeToString(checkSum[:]))

	exists := kvstore.IsUpdateExists(hash)
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
