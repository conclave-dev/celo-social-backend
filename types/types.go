package types

import "github.com/ethereum/go-ethereum/common"

type Profile struct {
	Address  common.Address `json:"address"`
	PhotoURL string         `json:"photoURL"`
}
