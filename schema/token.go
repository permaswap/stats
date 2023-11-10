package schema

import (
	"github.com/everFinance/everpay/token/utils"
)

type Token struct {
	ID        string // On Native-Chain tokenId; Special AR token: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA,0xcc9141efa8c20c7df0778748255b1487957811be"
	Symbol    string
	ChainType string // On everPay chainType; Special AR token: "arweave,ethereum"
	ChainID   string // On everPay chainId; Special AR token: "0,1"(mainnet) or "0,42"(testnet)
	Decimals  int    // On everPay decimals
}

func (t *Token) Tag() string {
	return utils.Tag(t.ChainType, t.Symbol, t.ID)
}
