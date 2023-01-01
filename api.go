package stats

import (
	"encoding/json"
	"net/http"
	"time"

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
		Tokens:          s.tokens,
		Pools:           s.pools,
		CurStats:        s.curStats,
	})
}

func (s *Stats) getStats(c *gin.Context) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	date := c.Query("date")
	if date == "" {
		c.JSON(http.StatusBadRequest, "err_no_param")
		return
	}
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
}
