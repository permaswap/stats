package stats

import (
	"github.com/everFinance/everpay/token"
	tokSchema "github.com/everFinance/everpay/token/schema"

	"github.com/permaswap/stats/schema"
)

func getPools(timestamp int64, chainID int64) (pools map[string]*schema.Pool) {
	// 2022-12-12
	if timestamp >= 1670774400000 {
		switch chainID {

		case 1: // everPay mainnet
			eth_usdc := &schema.Pool{
				TokenXTag: "ethereum-eth-0x0000000000000000000000000000000000000000",
				TokenYTag: "ethereum-usdc-0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
				FeeRatio:  schema.Fee003,
			}

			ar_usdc := &schema.Pool{
				TokenXTag: "arweave,ethereum-ar-AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA,0x4fadc7a98f2dc96510e42dd1a74141eeae0c1543",
				TokenYTag: "ethereum-usdc-0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
				FeeRatio:  schema.Fee003,
			}

			ar_eth := &schema.Pool{
				TokenXTag: "arweave,ethereum-ar-AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA,0x4fadc7a98f2dc96510e42dd1a74141eeae0c1543",
				TokenYTag: "ethereum-eth-0x0000000000000000000000000000000000000000",
				FeeRatio:  schema.Fee003,
			}

			usdc_usdt := &schema.Pool{
				TokenXTag: "ethereum-usdc-0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
				TokenYTag: "ethereum-usdt-0xdac17f958d2ee523a2206206994597c13d831ec7",
				FeeRatio:  schema.Fee0005,
			}

			pools = map[string]*schema.Pool{
				eth_usdc.ID():  eth_usdc,
				ar_eth.ID():    ar_eth,
				usdc_usdt.ID(): usdc_usdt,
				ar_usdc.ID():   ar_usdc,
			}

		case 5:
			tar_tusdc := &schema.Pool{
				TokenXTag: "bsc-tar-0xf1458ee7e9a2096bce7a21c160840a3a291bcb55",
				TokenYTag: "bsc-tusdc-0xf17a50ecc5fe5f476de2da5481cdd0f0ffef7712",
				FeeRatio:  schema.Fee003,
			}

			tar_tardrive := &schema.Pool{
				TokenXTag: "bsc-tar-0xf1458ee7e9a2096bce7a21c160840a3a291bcb55",
				TokenYTag: "bsc-tardrive-0xf4233b165f1b8da4f9aa94abc35c9ad2a7612979",
				FeeRatio:  schema.Fee003,
			}

			pools = map[string]*schema.Pool{
				tar_tardrive.ID(): tar_tardrive,
				tar_tusdc.ID():    tar_tusdc,
			}

		default:
			return nil
		}
		return
	} else {
		return nil
	}
}

func getTokens(timestamp int64, chainID int64) (tokens map[string]*token.Token) {
	tokenList := []*token.Token{}
	// 2022-12-12
	if timestamp >= 1670774400000 {
		switch chainID {
		case 1: // everPay mainnet{
			ar := token.New(
				"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA,0x4fadc7a98f2dc96510e42dd1a74141eeae0c1543",
				"AR", "arweave,ethereum", "0,1", 12, []tokSchema.TargetChain{},
			)
			eth := token.New(
				"0x0000000000000000000000000000000000000000", "ETH", "ethereum", "1", 18, []tokSchema.TargetChain{},
			)
			usdt := token.New(
				"0xdac17f958d2ee523a2206206994597c13d831ec7", "USDT", "ethereum", "1", 6, []tokSchema.TargetChain{},
			)
			usdc := token.New(
				"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "USDC", "ethereum", "1", 6, []tokSchema.TargetChain{},
			)
			tokenList = []*token.Token{ar, eth, usdc, usdt}

		case 5: // test network
			tar := token.New(
				"0xF1458EE7E9a2096BCE7a21c160840a3a291bcB55", "tAR", "bsc", "97", 12, []tokSchema.TargetChain{},
			)
			tusdc := token.New(
				"0xf17A50Ecc5Fe5f476DE2da5481cDD0F0ffef7712", "tUSDC", "bsc", "97", 6, []tokSchema.TargetChain{},
			)
			tardrive := token.New(
				"0xf4233B165F1b8DA4f9Aa94abC35C9ad2A7612979", "tARDRIVE", "bsc", "97", 18, []tokSchema.TargetChain{},
			)
			tokenList = []*token.Token{tar, tusdc, tardrive}
		default:
			return nil

		}

		tokens = map[string]*token.Token{}
		for _, t := range tokenList {
			tokens[t.Tag()] = t
		}

		return
	} else {
		return nil
	}
}
