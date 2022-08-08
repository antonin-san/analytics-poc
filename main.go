package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/30ea8288e72c4a5496e90faa2dc77c92")
	client, err := ethclient.Dial("ws://127.0.0.1:8545")
	if err != nil {
		fmt.Printf("%v\n", err)
		log.Fatal(err)
	}

	defer client.Close()

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		fmt.Printf("%v\n", err)
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			fmt.Printf("%v\n", err)
			log.Fatal(err)
		case header := <-headers:
			//fmt.Printf("%#v\n", header.Number)

			block, err := client.BlockByNumber(context.Background(), header.Number)

			if block == nil {
				//fmt.Print("nil block\n")
				continue
			}

			if err != nil {
				//fmt.Printf("BlockbyHash +%v \n", err)
				log.Fatal(err)
			}

			//fmt.Println("Block number: ", block.Number().Uint64())

			for _, tx := range block.Transactions() {
				//fmt.Println("Hash: ", tx.Hash().Hex())            // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
				//fmt.Println("Value: ", tx.Value().String())       // 10000000000000000
				//fmt.Println("Gas: ", tx.Gas())                    // 105000
				//fmt.Println("GasPrice: ", tx.GasPrice().Uint64()) // 102000000000
				//fmt.Println("Nonce: ", tx.Nonce())                // 110644
				//fmt.Println("Data: ", tx.Data())                  // []
				//fmt.Println("To: ", tx.To().Hex())                // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

				chainID, err := client.NetworkID(context.Background())
				if err != nil {
					//fmt.Printf("Network +%v \n", err)
					log.Fatal(err)
				}

				if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), nil); err == nil {
					fmt.Println("From: ", msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
					//fmt.Println("From: ")                   // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
				}

				receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
				if err != nil {
					//fmt.Printf("Tx Reiceip +%v \n", err)
					log.Fatal(err)
				}

				query := ethereum.FilterQuery{
					Addresses: []common.Address{*tx.To()},
					FromBlock: big.NewInt(int64(block.Number().Uint64())),
					ToBlock:   big.NewInt(int64(block.Number().Uint64())),
				}

				logs, err := client.FilterLogs(context.Background(), query)

				if err != nil {
					//fmt.Printf("Filte +%v \n", err)
					log.Fatal(err)
				}

				for i, _ := range logs {
					fmt.Println("Logs: ", i)
				}

				fmt.Println("Status end:", receipt.Status) // 1
			}
		}
	}
}
