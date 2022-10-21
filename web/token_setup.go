package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"sync"
)

var (
	tokensGiveawayMap = &tokensMap{
		m:  map[string]bool{},
		rw: &sync.RWMutex{},
	}
)

func setupEthClient(url string) (*ethclient.Client, error) {
	return ethclient.Dial(url)
}

func setupTokenContract(addr string, client *ethclient.Client) (*Token, error) {
	return NewToken(common.HexToAddress(addr), client)
}

// Execute the faucet token giveaway goroutine
func runTokenFaucet(token *Token) {
}
