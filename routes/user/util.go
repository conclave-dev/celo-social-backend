package user

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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

func getProfileFromHash(hash string) (kvstore.Profile, error) {
	p, err := kvstore.GetProfile(hash)
	if err != nil || p == "" {
		// Delete the user address if the claim hash is invalid
		return kvstore.Profile{}, fmt.Errorf("Error getting profile %s", err)
	}

	var profile kvstore.Profile
	err = json.Unmarshal([]byte(p), &profile)
	if err != nil {
		return kvstore.Profile{}, err
	}

	return profile, nil
}

func createNewProfile(address string) (types.User, error) {
	user, err := fetchAccount(address)
	if err != nil {
		return user, err
	}

	// Create new user profile and marshal into storable string type
	user.Profile = makeProfile(user)
	p, err := json.Marshal(user.Profile)
	if err != nil {
		return user, err
	}

	// Generate a hash from user profile and account data, and set as key to profile
	user.Hash = makeHash(user.Profile, user.AccountSummary, user.Metadata)
	_, err = kvstore.SetProfile(user.Hash, string(p))
	if err != nil {
		return user, err
	}

	_, err = kvstore.SetAddress(user.AccountSummary.Address.String(), user.Hash)
	if err != nil {
		return user, err
	}

	// Celo accounts are not required to have names, so we stop here if not set
	if user.AccountSummary.Name == "" {
		return user, nil
	}

	// Map username to address, to fetch hash if using username
	_, err = kvstore.SetUser(user.AccountSummary.Name, user.AccountSummary.Address.String())
	if err != nil {
		return user, err
	}

	return user, nil
}
