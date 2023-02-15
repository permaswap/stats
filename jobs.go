package stats

import (
	"encoding/json"

	"github.com/permaswap/stats/schema"
)

func (s *Stats) runJobs() {
	s.scheduler.Every(10).Minute().SingletonMode().Do(s.aggregateStats)
	s.scheduler.StartAsync()
}

func (s *Stats) aggregateStats() {
	statSnapshots, err := s.wdb.LoadAllStats()
	if err != nil {
		log.Error("failed to load all stats.", err, err)
	}
	pool := map[string]float64{}
	user := map[string]float64{}
	lp := map[string]map[string]float64{}
	lpReward := map[string]map[string]float64{}
	fee := float64(0)
	txCount := int64(0)
	for _, statSnapshot := range statSnapshots {
		b := []byte(statSnapshot.Stats)
		stats := schema.Stats{}
		if err := json.Unmarshal(b, &stats); err != nil {
			log.Error("failed to unmarshal stats snapshot")
			continue
		}
		fee += stats.Fee
		txCount += stats.TxCount
		for p, v := range stats.Pool {
			pool[p] += v
		}
		for u, v := range stats.User {
			user[u] += v
		}
		for accid, v := range stats.Lp {
			if _, ok := lp[accid]; ok {
				for poolID, v2 := range v {
					lp[accid][poolID] += v2
				}
			} else {
				lp[accid] = v
			}
		}
		for accid, r := range stats.LpReward {
			if _, ok := lpReward[accid]; ok {
				for poolID, v2 := range r {
					lpReward[accid][poolID] += v2
				}
			} else {
				lpReward[accid] = r
			}
		}
	}

	aggregate := schema.Aggregate{
		Stats: schema.Stats{
			Pool:     pool,
			User:     user,
			Lp:       lp,
			LpReward: lpReward,
			Fee:      fee,
			TxCount:  txCount,
		},
	}

	s.lockAggregate.Lock()
	defer s.lockAggregate.Unlock()
	s.aggregate = &aggregate
}
