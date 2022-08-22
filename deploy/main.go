package main

import (
	"context"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	arb "github.com/zombeer/go-arbitrage/gen"
)

func main() {
	keyphrase := os.Getenv("KEYPHRASE")
	ctx := context.Background()
	b, err := ioutil.ReadFile("./wallets/keystore_etherium")
	if err != nil {
		panic(err)
	}

	key, err := keystore.DecryptKey(b, keyphrase)
	println(key.Address.String())

	client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/13d446ced4df4fcba663dfed82716ad7")
	// // client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	nonce, err := client.PendingNonceAt(ctx, key.Address)
	if err != nil {
		panic(err)
	}
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		panic(err)
	}

	chainId, err := client.NetworkID(ctx)
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainId)
	if err != nil {
		panic(err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000)
	auth.Nonce = big.NewInt(int64(nonce))

	addr, _, _, err := arb.DeployArb(auth, client)
	if err != nil {
		panic(err)
	}

	println("Contract deployed...!")
	println("Contract address:", addr.String())
}
