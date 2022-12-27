package schema

type InfoRes struct {
	ChainID         int64   `json:"chainID"`
	RouterAddress   string  `json:"routerAddress"`
	StartAt         int64   `json:"startAt"`
	StartTxRawID    int64   `json:"startTxRawID"`
	StartTxEverHash string  `json:"startTxEverHash"`
	TotalUserVolume float64 `json:"totalUserVolume"`
}
