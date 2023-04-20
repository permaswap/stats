package stats

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/everFinance/everpay/account"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/permaswap/stats/schema"
)

func (s *Stats) runAPI(port string) {
	e := s.engine
	s.engine.Use(cors.Default())

	// api
	e.GET("/info", s.getInfo)
	e.GET("/stats", s.getStats)
	e.GET("/aggregate", s.getAggregate)

	s.server = &http.Server{
		Addr:    port,
		Handler: s.engine,
	}
	log.Info("api listening", "port", port)
	if err := s.server.ListenAndServe(); err != nil {
		log.Warn("server closed", "err", err)
	}
}

func (s *Stats) getInfo(c *gin.Context) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	c.JSON(http.StatusOK, schema.InfoRes{
		ChainID:         s.chainID,
		RouterAddress:   s.routerAddress,
		StartTxRawID:    s.startTxRawID,
		StartTxEverHash: s.startTxEverHash,
		//Tokens:          s.tokens,
		//Pools:           s.pools,
		CurStats: s.curStats,
	})
}

func (s *Stats) getStats(c *gin.Context) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	date := c.Query("date")
	if date != "" {
		t, _ := time.ParseInLocation("2006-01-02", date, time.Local)
		statSnapshot, err := s.wdb.FindStatsSnapshot(t)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		stats := schema.Stats{}
		b := []byte(statSnapshot.Stats)
		if err := json.Unmarshal(b, &stats); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, stats)
		return
	}

	startStr := c.DefaultQuery("start", "")
	if startStr != "" {
		c.JSON(http.StatusBadRequest, "err_no_param")
		return
	}
	start, err := time.ParseInLocation("2006-01-02", startStr, time.Local)
	if err != nil {
		c.JSON(http.StatusBadRequest, "err_invalid_param")
		return
	}
	endStr := c.DefaultQuery("end", "")
	if endStr != "" {
		c.JSON(http.StatusBadRequest, "err_no_param")
		return
	}
	end, err := time.ParseInLocation("2006-01-02", endStr, time.Local)
	if err != nil {
		c.JSON(http.StatusBadRequest, "err_invalid_param")
		return
	}
	diff := end.Sub(start)
	if diff.Hours() < 24 || diff.Hours() > 24*30 {
		c.JSON(http.StatusBadRequest, "err_invalid_param")
		return
	}

	statSnapshots, err := s.wdb.FindStatsSnapshots(start, end)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	stats := []schema.Stats{}
	for _, statSnapshot := range statSnapshots {
		s := schema.Stats{}
		b := []byte(statSnapshot.Stats)
		if err := json.Unmarshal(b, &s); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		stats = append(stats, s)
	}

	c.JSON(http.StatusOK, stats)
}

func (s *Stats) getAggregate(c *gin.Context) {
	s.lockAggregate.RLock()
	defer s.lockAggregate.RUnlock()

	if s.aggregate == nil {
		c.JSON(http.StatusOK, schema.AggregateRes{})
		return
	}

	aggregateRes := schema.AggregateRes{
		UserVolume: float64(0),
		LpVolume:   float64(0),
		LpReward:   float64(0),
	}
	for _, v := range s.aggregate.Stats.User {
		aggregateRes.UserVolume += v
	}
	for _, v := range s.aggregate.Stats.Lp {
		for _, v2 := range v {
			aggregateRes.LpVolume += v2
		}
	}
	for _, v := range s.aggregate.Stats.LpReward {
		for _, v2 := range v {
			aggregateRes.LpReward += v2
		}
	}

	if accid := c.Query("accid"); accid != "" {
		_, accid, err := account.IDCheck(accid)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		userAggregate := schema.UserAggregate{
			Address:    accid,
			LpVolume:   float64(0),
			LpReward:   float64(0),
			UserVolume: float64(0),
		}

		if v, ok := s.aggregate.Stats.User[accid]; ok {
			userAggregate.UserVolume = v
		}
		if v, ok := s.aggregate.Stats.Lp[accid]; ok {
			for _, v2 := range v {
				userAggregate.LpVolume += v2
			}
		}
		if v, ok := s.aggregate.Stats.LpReward[accid]; ok {
			for _, v2 := range v {
				userAggregate.LpReward += v2
			}
		}
		aggregateRes.User = userAggregate
	}

	c.JSON(http.StatusOK, aggregateRes)

}
