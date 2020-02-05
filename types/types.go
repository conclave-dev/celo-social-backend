package types

import "github.com/ethereum/go-ethereum/common"

type JSONResponse struct {
	Data interface{} `json:"data"`
}

type Profile struct {
	Address  common.Address `json:"address"`
	PhotoURL string         `json:"photoURL"`
}

type UpdateResponse struct {
	Hash   string
	Update string
}
