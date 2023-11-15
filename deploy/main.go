package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Lokal dosya sisteminden import

func main() {
	client, err := ethclient.Dial("https://goerli.infura.io/v3/4d57d898f5864252972f33b7811df35c")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Özel anahtar dosyasının adını ve yolunu belirtin
	privateKeyFilePath := filepath.Join("..", "private")

	// Özel anahtar dosyasını okuyun
	privateKeyBytes, err := os.ReadFile(privateKeyFilePath)
	if err != nil {
		log.Fatal(err)
	}

	privateKeyStr := string(privateKeyBytes)

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	fmt.Println(publicKey)
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	// }
	// add := crypto.PubkeyToAddress(*publicKeyECDSA)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	cAdd := common.HexToAddress("0x1d61531870D2CAf7deFAdfD5dbBe4AD053BfC087")
	// t, err := todo.NewTodo(cAdd, client)
	if err != nil {
		log.Fatal(cAdd)
	}

	tx, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	tx.GasLimit = 3000000
	tx.GasPrice = gasPrice

	// tra, err := t.Add(tx, "five Task")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(tra.Hash())

	// tasks, err := t.List(&bind.CallOpts{
	// 	From: add,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(tasks)

	// tra, err := t.Update(tx, big.NewInt(0), "update task content")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Toggle tx", tra.Hash())

	// tra, err := t.Remove(tx, big.NewInt(0))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Toggle tx", tra.Hash())

}
