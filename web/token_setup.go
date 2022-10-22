package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"sync"
)

var (
	tokensAmountToGive = big.NewInt(5)

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
func runTokenFaucet(token *Token, priv *privateKey) {
}

func executeTokenFaucetTick(client *ethclient.Client, priv *privateKey, token *Token) error {
	pending := tokensGiveawayMap.retrievePending()

	// Retrieve transaction opts
	opts, _, err := prepareContractWrite(client, 300000, priv)
	if err != nil {
		return err
	}

	// Send tokens to all the pending entries and update them
	for _, addr := range pending {
		token.Transfer(opts, common.HexToAddress(addr), tokensAmountToGive)

		// Mark entry as done
	}
	return nil
}

func prepareContractWrite(client *ethclient.Client, gasLimit uint64, privKey *privateKey) (*bind.TransactOpts, *big.Int, error) {
	// Retrieve chain id
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, nil, err
	}

	// Retrieve account nonce
	nonce, err := client.PendingNonceAt(context.Background(), privKey.From)
	if err != nil {
		return nil, nil, fmt.Errorf("Unable to get new nonce: %v", err)
	}

	// Get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, nil, fmt.Errorf("Unable to retrieve gas price: %v", err)
	}

	// Create auth
	auth, err := bind.NewKeyedTransactorWithChainID(privKey.Private, chainId)
	if err != nil {
		return nil, nil, fmt.Errorf("Unable to create transactor: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	return auth, chainId, nil
}
