package stats

import (
	"net/http"

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

}
