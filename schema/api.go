package schema

import (
	"github.com/everFinance/go-everpay/token"
)

type InfoRes struct {
	ChainID         int64                   `json:"chainID"`
	RouterAddress   string                  `json:"routerAddress"`
	StartTxRawID    int64                   `json:"startTxRawID"`
	StartTxEverHash string                  `json:"startTxEverHash"`
	Tokens          map[string]*token.Token `json:"tokens"`
	Pools           map[string]*Pool        `json:"pool"`
	CurStats        *Stats                  `json:"curStats"`
}
