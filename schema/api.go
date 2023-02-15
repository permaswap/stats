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

type UserAggregate struct {
	Address    string  `json:"address"`
	LpVolume   float64 `json:"lpVolume"`
	LpReward   float64 `json:"lpReward"`
	UserVolume float64 `json:"userVolume"`
}

type AggregateRes struct {
	LpVolume   float64 `json:"lpVolume"`
	LpReward   float64 `json:"lpReward"`
	UserVolume float64 `json:"userVolume"`
	User       UserAggregate
}
