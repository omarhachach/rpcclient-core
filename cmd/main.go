package main

import (
	"fmt"

	"github.com/omarhachach/rpcclient-core"
)

func main() {
	client, err := rpcclient.New(&rpcclient.Config{
		Host:                 "http://bitcoin.staging.dreamsescrow.com/",
		User:                 "bitcoind",
		Pass:                 "bitcoind",
		DisableAutoReconnect: false,
	})
	if err != nil {
		panic(err)
		return
	}

	count, err := client.GetBlockChainInfo()
	if err != nil {
		panic(err)
		return
	}

	fmt.Printf("Info: %#v\n", count)
}
