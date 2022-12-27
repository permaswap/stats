package stats

import (
	"fmt"

	"github.com/permaswap/stats/schema"
)

func InitPools(chainID int64) (pools map[string]*schema.Pool) {
	switch chainID {

	case 1: // everPay mainnet
		eth_usdc := &schema.Pool{
			"ethereum-eth-0x0000000000000000000000000000000000000000",
			"ethereum-usdc-0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
			schema.Fee003,
		}

		ar_usdc := &schema.Pool{
			"arweave,ethereum-ar-AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA,0x4fadc7a98f2dc96510e42dd1a74141eeae0c1543",
			"ethereum-usdc-0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
			schema.Fee003,
		}

		ar_eth := &schema.Pool{
			"arweave,ethereum-ar-AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA,0x4fadc7a98f2dc96510e42dd1a74141eeae0c1543",
			"ethereum-eth-0x0000000000000000000000000000000000000000",
			schema.Fee003,
		}

		usdc_usdt := &schema.Pool{
			"ethereum-usdc-0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
			"ethereum-usdt-0xdac17f958d2ee523a2206206994597c13d831ec7",
			schema.Fee0005,
		}

		pools = map[string]*schema.Pool{
			eth_usdc.ID():  eth_usdc,
			ar_eth.ID():    ar_eth,
			usdc_usdt.ID(): usdc_usdt,
			ar_usdc.ID():   ar_usdc,
			//ar_cfx.ID():    ar_cfx,
		}

	case 5:
		tar_tusdc := &schema.Pool{
			"bsc-tar-0xf1458ee7e9a2096bce7a21c160840a3a291bcb55",
			"bsc-tusdc-0xf17a50ecc5fe5f476de2da5481cdd0f0ffef7712",
			schema.Fee003,
		}

		tar_tardrive := &schema.Pool{
			"bsc-tar-0xf1458ee7e9a2096bce7a21c160840a3a291bcb55",
			"bsc-tardrive-0xf4233b165f1b8da4f9aa94abc35c9ad2a7612979",
			schema.Fee003,
		}

		pools = map[string]*schema.Pool{
			tar_tardrive.ID(): tar_tardrive,
			tar_tusdc.ID():    tar_tusdc,
		}

	default:
		panic(fmt.Sprintf("can not init pools, invalid chainID: %d\n", chainID))
	}

	return
}