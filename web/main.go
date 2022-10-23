package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type httpHandlerFunc func(*template.Template, http.ResponseWriter, *http.Request) error

func main() {
	tpl, err := template.New("faucet").ParseGlob("views/*")
	if err != nil {
		fmt.Printf("Unable to load views: %v\n", err)
		return
	}

	go startFaucet()

	http.HandleFunc("/", errorHandler(tpl, showFaucetPage))
	http.HandleFunc("/faucet", errorHandler(tpl, handleFaucetRequest))

	fmt.Println("Starting HTTP server in :8080")

	// Create custom server settings
	httpServer := &http.Server{
		Addr:              ":8080",
		Handler:           nil,
		ReadTimeout:       2 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       100 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	if err := httpServer.ListenAndServe(); err != nil {
		fmt.Printf("Unable to start HTTP server: %v\r\n", err)
	}
}

func errorHandler(tpl *template.Template, f httpHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if err := f(tpl, w, req); err != nil {
			w.Write([]byte(err.Error()))
		}
	}
}

const (
	nodeURL         = "http://localhost:8545"
	contractAddress = ""
	privateKeyStr   = ""
)

func startFaucet() {
	// Create ethclient
	client, err := setupEthClient(nodeURL)
	if err != nil {
		fmt.Printf("Unable to connect to node: %v\n", err)
		return
	}

	// Create contract object
	tokenContract, err := setupTokenContract(contractAddress, client)
	if err != nil {
		fmt.Printf("Unable to create contract object: %v\n", err)
		return
	}

	// Parse private key
	priv, err := parsePrivateKey(privateKeyStr)
	if err != nil {
		fmt.Printf("Unable to parse private key: %v\n", err)
		return
	}

	// Start Faucet
	runTokenFaucet(client, priv, tokenContract)
}
