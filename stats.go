package stats

import (
	"encoding/json"
	"errors"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"sync"
	"time"

	cacheSchema "github.com/everFinance/go-everpay/cache/schema"
	paySchema "github.com/everFinance/go-everpay/pay/schema"
	"github.com/everFinance/go-everpay/sdk"
	sdkSchema "github.com/everFinance/go-everpay/sdk/schema"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/permaswap/stats/logger"
	"github.com/permaswap/stats/schema"
)

var log = logger.New("stats")

type Stats struct {
	chainID         int64
	routerAddress   string
	startTxRawID    int64
	startTxEverHash string
	everClient      *sdk.Client
	engine          *gin.Engine
	server          *http.Server
	wdb             *WDB
	scheduler       *gocron.Scheduler

	sub *sdk.SubscribeTx
	//prevStats []*schema.Stats

	// lock to r/w curStats
	lock     sync.RWMutex
	curStats *schema.Stats

	lockAggregate sync.RWMutex
	aggregate     *schema.Aggregate
}

func New(chainID int64, routerAddr string, startTxRawID int64, startTxEverHash string, client *sdk.Client, dsn string) *Stats {
	return &Stats{
		chainID:         chainID,
		wdb:             NewWDB(dsn),
		engine:          gin.Default(),
		routerAddress:   routerAddr,
		startTxRawID:    startTxRawID,
		startTxEverHash: startTxEverHash,
		everClient:      client,
		scheduler:       gocron.NewScheduler(time.UTC),
	}
}

func (s *Stats) Run(port string) {
	s.wdb.Migrate()

	s.curStats = s.loadCurStats()
	startCursor := s.startTxRawID
	if s.curStats != nil && s.curStats.LastTxRawID > startCursor {
		startCursor = s.curStats.LastTxRawID
	}
	log.Info("running", "startTxRawID", startCursor, "curStats", s.curStats)

	s.sub = s.everClient.SubscribeTxs(sdkSchema.FilterQuery{
		StartCursor: startCursor,
		Address:     s.routerAddress,
		Action:      paySchema.TxActionBundle,
	})
	go func() {
		for tx := range s.sub.Subscribe() {
			s.processTx(tx)
		}
	}()

	go s.runAPI(port)

	s.aggregateStats()
	go s.runJobs()
}

func (s *Stats) Close() {
	s.sub.Unsubscribe()
}

func (s *Stats) findPool(pools map[string]*schema.Pool, x, y string) (poolID string) {
	for poolID, pool := range pools {
		if pool.TokenXTag == x && pool.TokenYTag == y {
			return poolID
		}
		if pool.TokenYTag == x && pool.TokenXTag == y {
			return poolID
		}
	}
	return
}

func (s *Stats) loadCurStats() (curStats *schema.Stats) {
	statSnapshot, err := s.wdb.LoadStats()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	if err != nil {
		log.Error("failed to load stats from db", "err", err)
		panic(err)
	}
	b := []byte(statSnapshot.Stats)
	if err := json.Unmarshal(b, &curStats); err != nil {
		log.Error("failed to unmarshal stats snapshot")
		panic(err)
	}
	return
}

func (s *Stats) processTx(tx cacheSchema.TxResponse) {

	t := time.Unix(tx.Nonce/1000, 0)
	log.Info("process new tx:", "RawId", tx.RawId, "EverHash", tx.EverHash)
	bundleData := paySchema.BundleData{}
	err := json.Unmarshal([]byte(tx.Data), &bundleData)
	if err != nil {
		log.Error("invalid router order data,ignore.", "tx data", tx.Data)
		return
	}

	internalStatus := cacheSchema.InternalStatus{}
	if err := json.Unmarshal([]byte(tx.InternalStatus), &internalStatus); err != nil {
		log.Error("failed to unmarshal tx internalStatus")
		return
	}

	if internalStatus.Status != cacheSchema.InternalStatusSuccess {
		log.Warn("tx failed, ignore", "tx.InternalStatus", tx.InternalStatus)
		return
	}

	poolStats, lpStats, lpRewardStats, userStats, feeStats := s.getStatsFromBundle(tx.Nonce, &bundleData.Bundle.Bundle)

	if poolStats == nil || lpStats == nil || lpRewardStats == nil || userStats == nil {
		log.Error("failed to get stats")
		return
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	if s.curStats == nil {
		log.Info("first tx")
		s.curStats = &schema.Stats{
			Date:            truncateToDay(t),
			StartTxRawID:    tx.RawId,
			StartTxEverHash: tx.EverHash,
			LastTxRawID:     tx.RawId,
			LastTxEverHash:  tx.EverHash,
			Pool:            poolStats,
			User:            userStats,
			Lp:              lpStats,
			LpReward:        lpRewardStats,
			Fee:             feeStats,
			TxCount:         1,
		}
	} else {
		curDate := getFormatedDate(s.curStats.Date)
		if curDate == getFormatedDate(t) {
			for i, v := range poolStats {
				s.curStats.Pool[i] += v
			}

			for i, v := range userStats {
				s.curStats.User[i] += v
			}

			for accid := range lpStats {
				if _, ok := s.curStats.Lp[accid]; ok {
					for poolID, v2 := range lpStats[accid] {
						s.curStats.Lp[accid][poolID] += v2
					}
				} else {
					s.curStats.Lp[accid] = lpStats[accid]
				}
			}

			for accid := range lpRewardStats {
				if _, ok := s.curStats.LpReward[accid]; ok {
					for poolID, r2 := range lpRewardStats[accid] {
						s.curStats.LpReward[accid][poolID] += r2
					}
				} else {
					s.curStats.LpReward[accid] = lpRewardStats[accid]
				}
			}

			s.curStats.LastTxRawID = tx.RawId
			s.curStats.LastTxEverHash = tx.EverHash
			s.curStats.Fee += feeStats
			s.curStats.TxCount += 1
		} else {
			log.Info("next day", "curDate", curDate, "tx date", getFormatedDate(t))
			if err := s.wdb.SaveStatsSnapshot(s.curStats, nil); err != nil {
				log.Error("Failed to save stats.", "stats date", s.curStats.Date)
			}
			s.curStats = &schema.Stats{
				Date:            truncateToDay(t),
				StartTxRawID:    tx.RawId,
				StartTxEverHash: tx.EverHash,
				LastTxRawID:     tx.RawId,
				LastTxEverHash:  tx.EverHash,
				Pool:            poolStats,
				User:            userStats,
				Lp:              lpStats,
				LpReward:        lpRewardStats,
				Fee:             feeStats,
				TxCount:         1,
			}
		}
	}
}

func (s *Stats) getStatsFromBundle(nonce int64, bundle *paySchema.Bundle) (
	poolStats map[string]float64,
	lpStats map[string]map[string]float64,
	lpRewardStats map[string]map[string]float64,
	userStats map[string]float64,
	feeStats float64) {

	pools := getPools(nonce, s.chainID)
	tokens := getTokens(nonce, s.chainID)

	poolStats = map[string]float64{}
	lpStats = map[string]map[string]float64{}
	lpRewardStats = map[string]map[string]float64{}
	userStats = map[string]float64{}
	feeStats = 0

	user := bundle.Items[0].From
	first := paySchema.BundleItem{}
	second := paySchema.BundleItem{}

	userInOut := map[string]*big.Int{}
	prices := map[string]float64{}

	zero := big.NewInt(0)
	for i, item := range bundle.Items {
		if i%2 == 0 {
			first = item
			continue
		}
		second = item

		poolID := s.findPool(pools, first.Tag, second.Tag)
		if poolID == "" {
			log.Error("failed to find pool", "x", first.Tag, "y", second.Tag)
			return nil, nil, nil, nil, 0
		}
		pool, ok := pools[poolID]
		if !ok {
			log.Error("failed to find the info of pool", "x", first.Tag, "y", second.Tag)
			return nil, nil, nil, nil, 0
		}

		_, ok = tokens[first.Tag]
		if !ok {
			log.Error("failed to find token", "tokenTag", first.Tag)
			return nil, nil, nil, nil, 0
		}
		_, ok = tokens[second.Tag]
		if !ok {
			log.Error("failed to find token", "tokenTag", first.Tag)
			return nil, nil, nil, nil, 0
		}

		lpAccID := first.From
		if lpAccID == user {
			lpAccID = first.To
		}

		var amount float64
		var decimals int

		symbol := tokens[first.Tag].Symbol
		price, ok := prices[symbol]
		if !ok {
			price = MustGetTokenPrice(symbol, "USDC", strconv.FormatInt(nonce, 10))
		}
		prices[symbol] = price

		amount, _ = strconv.ParseFloat(first.Amount, 64)
		decimals = tokens[first.Tag].Decimals
		volume := amount / math.Pow10(decimals) * price

		poolFeeRatio, _ := strconv.ParseFloat(pool.FeeRatio, 64)
		lpReward := volume * poolFeeRatio
		poolStats[poolID] += volume
		if _, ok := lpStats[lpAccID]; ok {
			lpStats[lpAccID][poolID] += volume
			lpRewardStats[lpAccID][poolID] += lpReward
		} else {
			lpStats[lpAccID] = map[string]float64{poolID: volume}
			lpRewardStats[lpAccID] = map[string]float64{poolID: lpReward}
		}

		if first.From == user {
			amount, _ := new(big.Int).SetString(first.Amount, 10)
			if v, ok := userInOut[first.Tag]; ok {
				userInOut[first.Tag] = new(big.Int).Sub(v, amount)
			} else {
				userInOut[first.Tag] = new(big.Int).Sub(zero, amount)
			}
			amount, _ = new(big.Int).SetString(second.Amount, 10)
			if v, ok := userInOut[second.Tag]; ok {
				userInOut[second.Tag] = new(big.Int).Add(v, amount)
			} else {
				userInOut[second.Tag] = amount
			}
		} else {
			amount, _ := new(big.Int).SetString(first.Amount, 10)
			if v, ok := userInOut[first.Tag]; ok {
				userInOut[first.Tag] = new(big.Int).Add(v, amount)
			} else {
				userInOut[first.Tag] = amount
			}
			amount, _ = new(big.Int).SetString(second.Amount, 10)
			if v, ok := userInOut[second.Tag]; ok {
				userInOut[second.Tag] = new(big.Int).Sub(v, amount)
			} else {
				userInOut[second.Tag] = new(big.Int).Sub(zero, amount)
			}
		}
	}

	tokenIn := ""
	amountIn := float64(0)
	for k, v := range userInOut {
		if v.Cmp(zero) == -1 {
			tokenIn = k
			amountIn, _ = new(big.Float).SetInt(v).Float64()
			amountIn *= -1
			break
		}
	}
	symbol := tokens[tokenIn].Symbol
	price, ok := prices[symbol]
	if !ok {
		price = MustGetTokenPrice(symbol, "USDC", strconv.FormatInt(nonce, 10))
	}
	prices[symbol] = price
	decimals := tokens[tokenIn].Decimals
	userStats[user] = amountIn / math.Pow10(decimals) * price

	if len(bundle.Items)%2 == 0 {
		return
	}

	feePath := bundle.Items[len(bundle.Items)-1]
	if tokens[tokenIn].Tag() != feePath.Tag {
		log.Error("invalid feepath token", "feepath tokenTag", feePath.Tag, "tokenIn", tokens[tokenIn].Tag())
		return
	}
	if amount, err := strconv.ParseFloat(feePath.Amount, 64); err == nil {
		feeStats = amount / math.Pow10(decimals) * price
		userStats[user] += feeStats
	}
	return
}
