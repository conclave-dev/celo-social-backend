package user

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/stella-zone/celo-social-backend/kvstore"
	_types "github.com/stella-zone/celo-social-backend/types"
	"github.com/stella-zone/celo-social-backend/util"
	"github.com/stella-zone/go-celo/types"
)

func handleUnclaimedUser(address string, w http.ResponseWriter) (_types.UserResponse, error) {
	var userResponse _types.UserResponse
	addressData, err := getAddressData(address, w)
	if err != nil {
		util.RespondWithError(err, w)
		return userResponse, err
	}

	userResponse.AccountSummary = addressData.AccountSummary
	userResponse.Metadata = addressData.Metadata

	if len(addressData.ClaimParams) != 0 {
		username := userResponse.AccountSummary.Name
		if username == "" {
			util.RespondWithError(errors.New(`Account is missing "name"`), w)
			return userResponse, err
		}

		userResponse.Profile = makeEmptyProfile(addressData.AccountSummary.Name)
		userResponse.Hash = generateUserHash(userResponse.Profile, userResponse.AccountSummary, userResponse.Metadata)

		u, err := json.Marshal(&_types.User{
			Hash: userResponse.Hash,
		})
		if err != nil {
			util.RespondWithError(err, w)
			return userResponse, err
		}
		kvstore.SetUser(username, string(u))

		p, err := json.Marshal(userResponse.Profile)
		if err != nil {
			util.RespondWithError(err, w)
			return userResponse, err
		}
		kvstore.SetProfile(userResponse.Hash, string(p))
	}

	return userResponse, nil
}

func getAddressData(address string, w http.ResponseWriter) (_types.AddressData, error) {
	var addressData _types.AddressData
	var accountSummary _types.AccountSummaryResponse
	err := fetchAccountSummary(address, &accountSummary, w)
	if err != nil {
		util.RespondWithError(err, w)
		return addressData, err
	}

	metadata, claimParams, err := fetchAndParseMetadata(accountSummary.Data.MetadataURL, w)
	if err != nil {
		util.RespondWithError(err, w)
		return addressData, err
	}

	addressData = _types.AddressData{
		AccountSummary: accountSummary.Data,
		Metadata:       metadata,
		ClaimParams:    claimParams,
	}

	return addressData, err
}

func fetchAccountSummary(address string, account *_types.AccountSummaryResponse, w http.ResponseWriter) error {
	url := fmt.Sprintf("%s/accounts/getAccountSummary", ApiURL)
	data := []byte(fmt.Sprintf(`{"address":"%s"}`, address))
	reqJSON := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", url, reqJSON)
	if err != nil {
		util.RespondWithError(err, w)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	d := json.NewDecoder(resp.Body)
	d.DisallowUnknownFields()
	err = d.Decode(&account)

	// An io.EOF error is returned by Decode() if the body is empty.
	if err != nil && !errors.Is(err, io.EOF) {
		util.HandleJSONDecodeError(err, w)
		return err
	}

	return nil
}

func fetchAndParseMetadata(metadataURL string, w http.ResponseWriter) (_types.Metadata, []string, error) {
	var metadata _types.Metadata

	m, err := util.SendGET(metadataURL)
	if err != nil {
		util.RespondWithError(err, w)
		return metadata, nil, err
	}

	err = json.Unmarshal(m, &metadata)
	if err != nil {
		util.RespondWithError(err, w)
		return metadata, nil, err
	}

	return metadata, getClaimParams(metadata.Claims), nil
}

func getClaimParams(claims _types.Claims) []string {
	http := fmt.Sprintf("http://%s", ClaimDomain)
	https := fmt.Sprintf("https://%s", ClaimDomain)
	var params []string

	for _, claim := range claims {
		if claim.Domain == "" {
			continue
		}

		switch true {
		case strings.Contains(claim.Domain, http):
			params = append(params, strings.ReplaceAll(claim.Domain, http, ""))
		case strings.Contains(claim.Domain, https):
			params = append(params, strings.ReplaceAll(claim.Domain, https, ""))
		default:
			break
		}
	}

	return params
}

func makeEmptyProfile(name string) _types.Profile {
	profile := _types.Profile{
		Name:        name,
		PhotoURL:    "",
		Email:       "",
		Description: "",
		Members:     []_types.Member{},
	}

	return profile
}

func generateUserHash(profile _types.Profile, accountSummary types.Account, metadata _types.Metadata) string {
	type UserHash struct {
		AccountSummary types.Account   `json:"accountSummary"`
		Profile        _types.Profile  `json:"profile"`
		Metadata       _types.Metadata `json:"metadata"`
	}

	data := fmt.Sprint(UserHash{
		Profile:  profile,
		Metadata: metadata,
	})

	// Convert decoded user object to sha256 checksum
	checkSum := sha256.Sum256([]byte(data))
	return fmt.Sprint(hex.EncodeToString(checkSum[:]))
}
