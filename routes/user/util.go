package user

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/stella-zone/celo-social-backend/types"
	celo "github.com/stella-zone/go-celo/types"
)

func makeProfile() types.Profile {
	profile := types.Profile{
		Name:    "",
		Photo:   "",
		Details: "",
		Website: "",
		Contact: types.Contact{
			Info: "",
			Type: "",
		},
		Members: []types.Member{},
	}

	return profile
}

func makeHash(profile types.Profile, accountSummary celo.Account, metadata types.Metadata) string {
	type UserHash struct {
		AccountSummary celo.Account   `json:"accountSummary"`
		Profile        types.Profile  `json:"profile"`
		Metadata       types.Metadata `json:"metadata"`
	}

	data := fmt.Sprint(UserHash{
		Profile:  profile,
		Metadata: metadata,
	})

	// Convert decoded user object to sha256 checksum
	checkSum := sha256.Sum256([]byte(data))
	return fmt.Sprint(hex.EncodeToString(checkSum[:]))
}
