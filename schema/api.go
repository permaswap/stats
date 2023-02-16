package schema

type InfoRes struct {
	ChainID         int64  `json:"chainID"`
	RouterAddress   string `json:"routerAddress"`
	StartTxRawID    int64  `json:"startTxRawID"`
	StartTxEverHash string `json:"startTxEverHash"`
	CurStats        *Stats `json:"curStats"`
}

type UserAggregate struct {
	Address    string  `json:"address"`
	LpVolume   float64 `json:"lpVolume"`
	LpReward   float64 `json:"lpReward"`
	UserVolume float64 `json:"userVolume"`
}

type AggregateRes struct {
	LpVolume   float64       `json:"lpVolume"`
	LpReward   float64       `json:"lpReward"`
	UserVolume float64       `json:"userVolume"`
	User       UserAggregate `json:"user"`
}
