package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTokenPriceByCoingecko(t *testing.T) {
	price, err := GetTokenPriceByCoingecko("arweave", "01-03-2023")
	assert.NoError(t, err)
	t.Log(price)
}
