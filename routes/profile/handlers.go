package profile

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/stella-zone/celo-social-backend/util"
)

func updateProfile(w http.ResponseWriter, r *http.Request) {
	var p Profile
	err := util.ParseRequestParameters(w, r, &p)
	if err != nil {
		util.RespondWithError(err, w)
		return
	}

	// Convert decoded request body to sha256 checksum
	rawHash := sha256.Sum256([]byte(fmt.Sprint(p)))
	hash := fmt.Sprint(hex.EncodeToString(rawHash[:]))

	w.Write([]byte(hash))
}
