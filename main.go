package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"github.com/ava-labs/coreth/accounts/keystore"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zombeer/go-arbitrage/entities"
	arb "github.com/zombeer/go-arbitrage/gen"
	"github.com/zombeer/go-arbitrage/ierc20"
)

const ARB_CONTRACT_ADDRESS = "0x39637849591d7075f80fd0df4aa4c403612457b2"

func calcDiffPercentage(start, end *big.Int) float64 {
	x1 := end.Int64() - start.Int64()
	result := float64(x1) / float64(start.Int64()) * 100
	return result
}

func main() {
	ctx := context.Background()
	keyphrase := os.Getenv("KEYPHRASE")
	b, err := ioutil.ReadFile("./wallets/keystore_etherium")
	if err != nil {
		panic(err)
	}
	key, err := keystore.DecryptKey(b, keyphrase)
	if err != nil {
		log.Fatalf("key decoding error: %v", err)
	}

	// client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/13d446ced4df4fcba663dfed82716ad7")
	client, err := ethclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to node..")
	defer client.Close()

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err = <-sub.Err():
			log.Fatal(err)
		case <-headers:

			bn, err := client.BlockNumber(ctx)
			fmt.Printf("============ Block # %d =============", bn)
			fmt.Println()
			if err != nil {
				panic(err)
			}

			arbContractAddress := common.HexToAddress(ARB_CONTRACT_ADDRESS)
			arbBot, err := arb.NewArb(arbContractAddress, client)
			if err != nil {
				panic(err)
			}

			callOpts := bind.CallOpts{
				BlockNumber: big.NewInt(int64(bn)),
				Pending:     false,
				Context:     ctx,
			}

			weth := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
			wethToken, err := ierc20.NewIerc20(weth, client)
			if err != nil {
				log.Fatal()
			}

			wethBalance, err := wethToken.BalanceOf(&callOpts, arbContractAddress)
			if err != nil {
				log.Fatal(err)
			}

			for _, args := range entities.GetDualDexPermutations() {
				nonce, err := client.PendingNonceAt(ctx, key.Address)
				if err != nil {
					panic(err)
				}
				gasPrice, err := client.SuggestGasPrice(ctx)
				if err != nil {
					panic(err)
				}
				gasPrice = big.NewInt(gasPrice.Int64() + 1000000000)

				chainId, err := client.NetworkID(ctx)
				if err != nil {
					panic(err)
				}

				auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainId)
				if err != nil {
					panic(err)
				}

				auth.GasPrice = gasPrice
				auth.GasLimit = uint64(300000)
				auth.Nonce = big.NewInt(int64(nonce))

				maxGasFee := auth.GasLimit * gasPrice.Uint64()
				result := EstimateDualDex(arbBot, wethBalance, bn, args[0], args[1], args[2], args[3]).Uint64() - maxGasFee
				percentage := calcDiffPercentage(wethBalance, big.NewInt(int64(result)))

				fmt.Printf("%s <-> %s[%s/%s]: ", args[0].Name, args[1].Name, args[2].Name, args[3].Name)
				fmt.Printf("Result: %.2f%%", percentage)
				// fmt.Println("Contract WETH Balance:", wethBalance.String())

				if percentage > 0 {
					println("It seems we've found something... trading..")

					_, err = arbBot.DualDexTrade(
						auth,
						args[0].CommonAddress(),
						args[1].CommonAddress(),
						args[2].CommonAddress(),
						args[3].CommonAddress(),
						wethBalance)
					if err != nil {
						panic(err)
					}
				}
			}
		}
	}
}

func EstimateDualDex(arbBot *arb.Arb, amount *big.Int, bn uint64, r1, r2, tkn1, tkn2 entities.Entity) *big.Int {
	callOpts := bind.CallOpts{
		BlockNumber: big.NewInt(int64(bn)),
		Pending:     false,
		Context:     context.Background(),
	}
	result, err := arbBot.EstimateDualDexTrade(
		&callOpts,
		r1.CommonAddress(),
		r2.CommonAddress(),
		tkn1.CommonAddress(),
		tkn2.CommonAddress(),
		amount,
	)
	if err != nil {
		log.Fatalf("error in estimating: %v", err.Error())
	}
	return result
}
