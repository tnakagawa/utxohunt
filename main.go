package main

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
)

var (
	// we're running on testnet3
	testnet3Parameters = &chaincfg.TestNet3Params
)

// This is the main function -- c

func main() {
	fmt.Printf("6.892 class program - hw1 - utxohunt\n")

	// Task #1 make an address pair
	// Call AddressFrom PrivateKey() to make a keypair

	//	s, err := AddressFromPrivateKey()
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("address: %s\n", s)

	// Task #2 make a transaction
	// Call EZTxBuilder to make a transaction
	//	tx := EZTxBuilder()
	//	var buf bytes.Buffer
	//	tx.Serialize(&buf)
	//	fmt.Printf("tx in hex:\n%x\n", buf.Bytes())
	//

	// task 3, call OpReturnTxBuilder() the same way EZTxBuilder() was used

	return
}
