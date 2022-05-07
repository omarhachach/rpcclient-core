package main

import (
	"fmt"

	"github.com/omarhachach/rpcclient-core"
)

func main() {
	// TODO: build cli
	client, err := rpcclient.New(&rpcclient.Config{
		Host:                 "http://localhost:8333",
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
