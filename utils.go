package stats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson"
	"gopkg.in/h2non/gentleman.v2"
)

func GetTokenName(symbol string) (name string) {
	name = ""
	symbolToName := map[string]string{
		"USDC": "usdc",
		"USDT": "usdt",
		"AR":   "arweave",
		"ETH":  "ethereum",
	}
	name = symbolToName[symbol]
	return name
}
func GetTokenPriceByRedstone(tokenSymbol string, currency string, timestamp string) (float64, error) {
	cli := gentleman.New()
	cli.URL("https://api.redstone.finance")
	req := cli.Request()
	req.AddPath("/prices")
	req.AddQuery("symbols", fmt.Sprintf("%s,%s", strings.ToUpper(tokenSymbol), strings.ToUpper(currency)))
	req.AddQuery("provider", "redstone")
	if timestamp != "" {
		req.AddQuery("toTimestamp", timestamp)
	}

	resp, err := req.Send()
	if err != nil {
		return 0.0, err
	}

	if !resp.Ok {
		return 0.0, fmt.Errorf("get token: %s currency: %s prices from redstone failed", tokenSymbol, currency)
	}
	defer resp.Close()
	tokenJsonPath := fmt.Sprintf("%s.value", strings.ToUpper(tokenSymbol))
	currencyJsonPath := fmt.Sprintf("%s.value", strings.ToUpper(currency))
	prices := gjson.GetManyBytes(resp.Bytes(), tokenJsonPath, currencyJsonPath)
	if len(prices) != 2 {
		return 0.0, fmt.Errorf("get token: %s currency: %s prices from redstone failed, response price number incorrect", tokenSymbol, currency)
	}
	tokenPrice := prices[0].Float()
	currencyPrice := prices[1].Float()
	if currencyPrice <= 0.0 {
		return 0.0, fmt.Errorf("get currency: %s price from redstone less than 0.0; currencyPrice: %f", currency, currencyPrice)
	}
	price := tokenPrice / currencyPrice
	return price, nil
}

type CoingeckoData struct {
	MarketData struct {
		CurrentPrice struct {
			USD float64 `json:"usd"`
		} `json:"current_price"`
	} `json:"market_data"`
}

func GetTokenPriceByCoingecko(token, date string) (float64, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s/history?date=%s", token, date)

	resp, err := http.Get(url)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0.0, err
	}

	var data CoingeckoData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0.0, err
	}
	return data.MarketData.CurrentPrice.USD, nil
}

func MustGetTokenPriceByRedstone(tokenSymbol string, currency string, timestamp string) float64 {
	for {
		price, err := GetTokenPriceByRedstone(tokenSymbol, "USDC", timestamp)
		if err == nil {
			return price
		}
		log.Warn("failed to get price from redstone", "err", err, "tokenSymbol", tokenSymbol, "currency", currency, "timestamp", timestamp)

		time.Sleep(1 * time.Second)
	}
}

func MustGetTokenPrice(symbol string, currency string, timestamp string) (price float64) {
	if symbol == "tUSDC" {
		price = MustGetTokenPriceByRedstone("USDC", "USDC", timestamp)
	} else if symbol == "tAR" {
		price = MustGetTokenPriceByRedstone("AR", "USDC", timestamp)
	} else if symbol == "tARDRIVE" {
		price = 3.5
	} else {
		price = MustGetTokenPriceByRedstone(symbol, "USDC", timestamp)
	}
	return price
}

func GetTokenPrice(tokenSymbol string, currency string, timestamp string, date string) (price float64) {
	var err error
	price = 0.0
	if tokenSymbol == "tUSDC" {
		price = 1.0
	} else if tokenSymbol == "tAR" {
		price = 10.0
	} else if tokenSymbol == "tARDRIVE" {
		price = 3.5
	} else {
		price, err = GetTokenPriceByRedstone(tokenSymbol, "USDC", timestamp)
		if err != nil {
			tokenName := GetTokenName(tokenSymbol)
			if tokenName != "" {
				price, _ = GetTokenPriceByCoingecko(tokenName, date)
			}
		}
	}
	return price
}

func getFormatedDate(t time.Time) string {
	return fmt.Sprintf("%s-%s-%s", strconv.Itoa(t.Year()), t.Month().String(), strconv.Itoa(t.Day()))
}

func getFormatedDate2(t time.Time) string {
	return fmt.Sprintf("%s-%s-%s", strconv.Itoa(t.Day()), t.Month().String(), strconv.Itoa(t.Year()))
}

func getFormatedDateTime(t time.Time) string {
	return fmt.Sprintf("%s-%s-%s %d:%d:%d", strconv.Itoa(t.Year()), t.Month().String(), strconv.Itoa(t.Day()),
		t.Hour(), t.Minute(), t.Second())
}

func truncateToDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
