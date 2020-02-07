package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	_types "github.com/stella-zone/celo-social-backend/types"
	"github.com/stella-zone/celo-social-backend/util"
)

func getAddressAccountSummaryAndClaims(address common.Address, w http.ResponseWriter) (_types.AccountSummaryAndClaims, error) {
	var accountSummaryAndClaims _types.AccountSummaryAndClaims
	var accountSummary _types.AccountSummaryResponse
	err := fetchAccountSummary(address, &accountSummary, w)
	if err != nil {
		util.RespondWithError(err, w)
		return accountSummaryAndClaims, err
	}

	claims, err := fetchAndParseMetadata(accountSummary.Data.MetadataURL, w)
	if err != nil {
		util.RespondWithError(err, w)
		return accountSummaryAndClaims, err
	}

	accountSummaryAndClaims = _types.AccountSummaryAndClaims{
		AccountSummary: accountSummary.Data,
		Claims:         claims,
	}

	return accountSummaryAndClaims, err
}

func fetchAccountSummary(address common.Address, account *_types.AccountSummaryResponse, w http.ResponseWriter) error {
	url := fmt.Sprintf("%s/accounts/getAccountSummary", ApiURL)
	data := []byte(fmt.Sprintf(`{"address":"%s"}`, address.String()))
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

func fetchAndParseMetadata(metadataURL string, w http.ResponseWriter) (_types.Claims, error) {
	m, err := util.SendGET(metadataURL)
	if err != nil {
		util.RespondWithError(err, w)
		return nil, err
	}

	var metadata _types.Metadata
	err = json.Unmarshal(m, &metadata)
	if err != nil {
		util.RespondWithError(err, w)
		return nil, err
	}

	var claims _types.Claims
	for _, claim := range metadata.Claims {
		if claim.Domain == "" {
			continue
		}

		// Check if the metadata's claim domain is for Celo Social
		if strings.Contains(claim.Domain, ClaimDomain) {
			claims = append(claims, claim)
		}
	}

	return claims, nil
}
