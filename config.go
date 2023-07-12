package stats

import (
	"github.com/everFinance/everpay/token"
	tokSchema "github.com/everFinance/everpay/token/schema"

	"github.com/permaswap/stats/schema"
)

// mainnet
var ar = token.New(
	"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA,0x4fadc7a98f2dc96510e42dd1a74141eeae0c1543",
	"AR", "arweave,ethereum", "0,1", 12, []tokSchema.TargetChain{},
)
var eth = token.New(
	"0x0000000000000000000000000000000000000000", "ETH", "ethereum", "1", 18, []tokSchema.TargetChain{},
)
var usdt = token.New(
	"0xdac17f958d2ee523a2206206994597c13d831ec7", "USDT", "ethereum", "1", 6, []tokSchema.TargetChain{},
)
var usdc = token.New(
	"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "USDC", "ethereum", "1", 6, []tokSchema.TargetChain{},
)
var ardrive = token.New(
	"-8A6RexFkpfWwuyVO98wzSFZh0d6VJuI-buTJvlwOJQ", "ARDRIVE", "arweave", "0", 18, []tokSchema.TargetChain{},
)

var acnh = token.New(
	"0x72247989079dA354c9f0a6886B965bcc86550F8a", "ACNH", "everpay", "1", 8, []tokSchema.TargetChain{},
)

var ans = token.New("0x937EFa4a5Ff9d65785691b70a1136aAf8aDA7e62", "ANS", "ethereum", "1", 18, []tokSchema.TargetChain{})

var u = token.New("KTzTXT_ANmF84fWEKHzWURD1LWd9QaFR9yfYUwH2Lxw", "U", "arweave", "0", 6, []tokSchema.TargetChain{})
var stamp = token.New("TlqASNDLA1Uh8yFiH-BzR_1FDag4s735F3PoUFEv2Mo", "STAMP", "arweave", "0", 6, []tokSchema.TargetChain{})

var eth_usdc = &schema.Pool{
	TokenXTag: "ethereum-eth-0x0000000000000000000000000000000000000000",
	TokenYTag: "ethereum-usdc-0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
	FeeRatio:  schema.Fee003,
}

var ar_usdc = &schema.Pool{
	TokenXTag: "arweave,ethereum-ar-AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA,0x4fadc7a98f2dc96510e42dd1a74141eeae0c1543",
	TokenYTag: "ethereum-usdc-0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
	FeeRatio:  schema.Fee003,
}

var ar_eth = &schema.Pool{
	TokenXTag: "arweave,ethereum-ar-AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA,0x4fadc7a98f2dc96510e42dd1a74141eeae0c1543",
	TokenYTag: "ethereum-eth-0x0000000000000000000000000000000000000000",
	FeeRatio:  schema.Fee003,
}

var usdc_usdt = &schema.Pool{
	TokenXTag: "ethereum-usdc-0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
	TokenYTag: "ethereum-usdt-0xdac17f958d2ee523a2206206994597c13d831ec7",
	FeeRatio:  schema.Fee0005,
}

var eth_usdt = &schema.Pool{
	TokenXTag: "ethereum-eth-0x0000000000000000000000000000000000000000",
	TokenYTag: "ethereum-usdt-0xdac17f958d2ee523a2206206994597c13d831ec7",
	FeeRatio:  schema.Fee003,
}

var ar_usdt = &schema.Pool{
	TokenXTag: "arweave,ethereum-ar-AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA,0x4fadc7a98f2dc96510e42dd1a74141eeae0c1543",
	TokenYTag: "ethereum-usdt-0xdac17f958d2ee523a2206206994597c13d831ec7",
	FeeRatio:  schema.Fee003,
}

var ar_ardrive = &schema.Pool{
	TokenXTag: "arweave,ethereum-ar-AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA,0x4fadc7a98f2dc96510e42dd1a74141eeae0c1543",
	TokenYTag: "arweave-ardrive--8A6RexFkpfWwuyVO98wzSFZh0d6VJuI-buTJvlwOJQ",
	FeeRatio:  schema.Fee003,
}

var usdc_acnh = &schema.Pool{
	TokenXTag: "ethereum-usdc-0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
	TokenYTag: "everpay-acnh-0x72247989079da354c9f0a6886b965bcc86550f8a",
	FeeRatio:  schema.Fee0005,
}

var ar_ans = &schema.Pool{
	TokenXTag: "arweave,ethereum-ar-AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA,0x4fadc7a98f2dc96510e42dd1a74141eeae0c1543",
	TokenYTag: "ethereum-ans-0x937efa4a5ff9d65785691b70a1136aaf8ada7e62",
	FeeRatio:  schema.Fee003,
}

var ar_u = &schema.Pool{
	TokenXTag: "arweave,ethereum-ar-AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA,0x4fadc7a98f2dc96510e42dd1a74141eeae0c1543",
	TokenYTag: "arweave-u-KTzTXT_ANmF84fWEKHzWURD1LWd9QaFR9yfYUwH2Lxw",
	FeeRatio:  schema.Fee003,
}

var ar_stamp = &schema.Pool{
	TokenXTag: "arweave,ethereum-ar-AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA,0x4fadc7a98f2dc96510e42dd1a74141eeae0c1543",
	TokenYTag: "arweave-stamp-TlqASNDLA1Uh8yFiH-BzR_1FDag4s735F3PoUFEv2Mo",
	FeeRatio:  schema.Fee003,
}

// testnet
var tar_tt = token.New(
	"0xF1458EE7E9a2096BCE7a21c160840a3a291bcB55", "tAR", "bsc", "97", 12, []tokSchema.TargetChain{},
)
var tusdc_tt = token.New(
	"0xf17A50Ecc5Fe5f476DE2da5481cDD0F0ffef7712", "tUSDC", "bsc", "97", 6, []tokSchema.TargetChain{},
)
var tardrive_tt = token.New(
	"0xf4233B165F1b8DA4f9Aa94abC35C9ad2A7612979", "tARDRIVE", "bsc", "97", 18, []tokSchema.TargetChain{},
)
var acnh_tt = token.New(
	"0x0000000000000000000000000000000000000003", "ACNH", "everpay", "5", 6, []tokSchema.TargetChain{},
)

var tar_tusdc_tt = &schema.Pool{
	TokenXTag: "bsc-tar-0xf1458ee7e9a2096bce7a21c160840a3a291bcb55",
	TokenYTag: "bsc-tusdc-0xf17a50ecc5fe5f476de2da5481cdd0f0ffef7712",
	FeeRatio:  schema.Fee003,
}

var tar_tardrive_tt = &schema.Pool{
	TokenXTag: "bsc-tar-0xf1458ee7e9a2096bce7a21c160840a3a291bcb55",
	TokenYTag: "bsc-tardrive-0xf4233b165f1b8da4f9aa94abc35c9ad2a7612979",
	FeeRatio:  schema.Fee003,
}

var tardrive_tusdc_tt = &schema.Pool{
	TokenXTag: "bsc-tardrive-0xf4233b165f1b8da4f9aa94abc35c9ad2a7612979",
	TokenYTag: "bsc-tusdc-0xf17a50ecc5fe5f476de2da5481cdd0f0ffef7712",
	FeeRatio:  schema.Fee003,
}

var acnh_tusdc_tt = &schema.Pool{
	TokenXTag: "everpay-acnh-0x0000000000000000000000000000000000000003",
	TokenYTag: "bsc-tusdc-0xf17a50ecc5fe5f476de2da5481cdd0f0ffef7712",
	FeeRatio:  schema.Fee003,
}

func getPools(timestamp int64, chainID int64) (pools map[string]*schema.Pool) {
	// 2022-12-12
	if timestamp >= 1670774400000 {
		switch chainID {
		// everPay mainnet
		case 1:
			pools = map[string]*schema.Pool{
				eth_usdc.ID():   eth_usdc,
				ar_eth.ID():     ar_eth,
				usdc_usdt.ID():  usdc_usdt,
				ar_usdc.ID():    ar_usdc,
				ar_usdt.ID():    ar_usdt,
				eth_usdt.ID():   eth_usdt,
				ar_ardrive.ID(): ar_ardrive,
				usdc_acnh.ID():  usdc_acnh,
				ar_ans.ID():     ar_ans,
				ar_u.ID():       ar_u,
				ar_stamp.ID():   ar_stamp,
			}

		case 5:
			pools = map[string]*schema.Pool{
				tar_tardrive_tt.ID():   tar_tardrive_tt,
				tar_tusdc_tt.ID():      tar_tusdc_tt,
				tardrive_tusdc_tt.ID(): tardrive_tusdc_tt,
				acnh_tusdc_tt.ID():     acnh_tusdc_tt,
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
			tokenList = []*token.Token{ar, eth, usdc, usdt, ardrive, acnh, ans, u, stamp}
		case 5: // test network
			tokenList = []*token.Token{tar_tt, tusdc_tt, tardrive_tt, acnh_tt}
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
