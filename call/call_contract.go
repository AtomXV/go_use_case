package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	// "github.com/gin-gonic/gin"
)

func main() {
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("f52db4ce18a2e3b07607c51a6ed2c148bfbf72f3e5d829561d96672872f52478")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyEDCSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("公钥类型错误")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyEDCSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(600000)
	auth.Value = big.NewInt(0)

	address := common.HexToAddress("0x15c16bA0c72B78264eb3D3A0E160C7D8695730a1")
	instance, err := NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("contract is loaded")
	// _ = instance

	key := [32]byte{}
	value := [32]byte{}

	copy(key[:], []byte("foo"))
	copy(value[:], []byte("fighters"))

	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tx.Hash().Hex())

	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(result[:]))
}
