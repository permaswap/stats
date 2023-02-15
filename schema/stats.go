package schema

import "time"

type Stats struct {
	Date            time.Time                     `json:"date"`
	StartTxRawID    int64                         `json:"startTxRawID"`
	StartTxEverHash string                        `json:"startTxEverHash"`
	LastTxRawID     int64                         `json:"lastTxRawID"`
	LastTxEverHash  string                        `json:"lastTxEverHash"`
	Pool            map[string]float64            `json:"pool"`     // poolid -> volume in usd
	User            map[string]float64            `json:"user"`     // accid -> volume in usd
	Lp              map[string]map[string]float64 `json:"lp"`       // lp accid -> poolid -> volume in usd
	LpReward        map[string]map[string]float64 `json:"lpReward"` // lp accid -> poolid -> lp reward in usd
	Fee             float64                       `json:"fee"`      // fee in usd
	TxCount         int64                         `json:"txCount"`  // txs count
}

type Aggregate struct {
	Start time.Time `json:"stat"`
	End   time.Time `json:"end"`
	Stats Stats     `json:"stats"`
}

func (s *Stats) TotalUserVolume() (total float64) {
	for _, v := range s.User {
		total += v
	}
	return
}
