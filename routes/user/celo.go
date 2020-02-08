package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/stella-zone/celo-social-backend/types"
	"github.com/stella-zone/celo-social-backend/util"
	celo "github.com/stella-zone/go-celo/types"
)

func fetchAccount(address string) (types.User, error) {
	var user types.User
	accountSummary, err := fetchAccountSummary(address)
	if err != nil {
		return user, err
	}

	metadata, err := fetchMetadata(accountSummary.MetadataURL)
	if err != nil {
		return user, err
	}

	user = types.User{
		AccountSummary: accountSummary,
		Metadata:       metadata,
	}
	return user, nil
}

func fetchAccountSummary(address string) (celo.Account, error) {
	var accountSummaryResponse struct {
		Data celo.Account `json:"data"`
	}
	url := fmt.Sprintf("%s/accounts/getAccountSummary", CeloAPI)
	data := []byte(fmt.Sprintf(`{"address":"%s"}`, address))
	reqJSON := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", url, reqJSON)
	if err != nil {
		return accountSummaryResponse.Data, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return accountSummaryResponse.Data, err
	}
	defer resp.Body.Close()

	d := json.NewDecoder(resp.Body)
	d.DisallowUnknownFields()

	err = d.Decode(&accountSummaryResponse)

	// An io.EOF error is returned by Decode() if the body is empty.
	if err != nil && !errors.Is(err, io.EOF) {
		return accountSummaryResponse.Data, util.HandleJSONDecodeError(err)
	}

	return accountSummaryResponse.Data, nil
}

func fetchMetadata(metadataURL string) (types.Metadata, error) {
	var metadata types.Metadata

	m, err := util.SendGET(metadataURL)
	if err != nil {
		return metadata, err
	}

	err = json.Unmarshal(m, &metadata)
	if err != nil {
		return metadata, err
	}

	return metadata, nil
}

func findSocialClaimHash(claims types.Claims) string {
	// Find newest social claim by iterating backwards
	for i := len(claims) - 1; i >= 0; i = i - 1 {
		if claims[i].Domain == "" {
			continue
		}

		// Identify the social claim and remove the hash (should always be 3rd if claim domain is correct)
		// Example: https://celo.social/c3b28788238d0f4b0068e9a1613c1192b69af66e6109cf9042ad0d14e945d9df
		if strings.Contains(claims[i].Domain, "https://celo.social") {
			return strings.Split(claims[i].Domain, "/")[3]
		}
	}

	return ""
}
