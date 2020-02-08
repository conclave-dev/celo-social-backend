package user

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/stella-zone/celo-social-backend/kvstore"
	"github.com/stella-zone/celo-social-backend/types"
	celo "github.com/stella-zone/go-celo/types"
)

func makeProfile(user types.User) kvstore.Profile {
	profile := kvstore.Profile{
		Name:    user.AccountSummary.Name,
		Address: user.AccountSummary.Address.String(),
		Photo:   "",
		Details: "",
		Website: "",
		Contact: kvstore.Contact{
			Info: "",
			Type: "",
		},
		Members: []kvstore.Member{},
	}

	return profile
}

func makeHash(profile kvstore.Profile, accountSummary celo.Account, metadata types.Metadata) string {
	type UserHash struct {
		AccountSummary celo.Account    `json:"accountSummary"`
		Profile        kvstore.Profile `json:"profile"`
		Metadata       types.Metadata  `json:"metadata"`
	}

	data := fmt.Sprint(UserHash{
		Profile:  profile,
		Metadata: metadata,
	})

	// Convert decoded user object to sha256 checksum
	checkSum := sha256.Sum256([]byte(data))
	return fmt.Sprint(hex.EncodeToString(checkSum[:]))
}
