package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/everFinance/go-everpay/sdk"
	"github.com/permaswap/stats"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "auction",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "port", Value: ":8080", EnvVars: []string{"PORT"}},
			&cli.StringFlag{Name: "pay", Value: "https://api.everpay.io", Usage: "pay url", EnvVars: []string{"PAY"}},
			&cli.StringFlag{Name: "router", Value: "0xd110107adb30bce6c0646eaf77cc1c815012331d", Usage: "router address", EnvVars: []string{"ROUTER"}},
			&cli.Int64Flag{Name: "ever_chain_id", Value: 1, Usage: "ever chain chainId", EnvVars: []string{"EVER_CHAIN_ID"}},
			&cli.Int64Flag{Name: "start_tx_rawid", Value: 342195, Usage: "rawid of start tx", EnvVars: []string{"START_TX_RAWID"}},
			&cli.StringFlag{Name: "start_tx_everhash", Value: "0x3dc15dc19b9a817d439685185f2fbec18938b4295c7a3d2a2ad0655313c732f1", Usage: "everhash start tx", EnvVars: []string{"START_TX_EVERHASH"}},
			&cli.StringFlag{Name: "mysql", Value: "root@tcp(127.0.0.1:3306)/stats?charset=utf8mb4&parseTime=True&loc=Local", Usage: "mysql dsn", EnvVars: []string{"MYSQL"}},
		},
		Action: run,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	client := sdk.NewClient(c.String("pay"))

	s := stats.New(c.Int64("ever_chain_id"), c.String("router"), c.Int64("start_tx_rawid"), c.String("start_tx_everhash"), client, c.String("mysql"))
	s.Run(c.String("port"))

	<-signals
	s.Close()

	return nil
}
