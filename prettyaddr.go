package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

var numZeros int
var numThreads int

func init()  {
	flag.IntVar(&numZeros, "num-zeros", 1, "Number of zero bytes at the end of address")
	flag.IntVar(&numThreads, "num-threads", 4, "Number of threads to mine")
}
func main() {
	done := make(chan interface{})
	for i := 0; i < numThreads; i++ {
		go genKey(numZeros, done)
	}
	<-done
}

func genKey(numZeros int, done chan<- interface{}) {
	bs := make([]byte, numZeros)
	for {
		sk, err := crypto.GenerateKey()
		if err != nil {
			log.Panic(err.Error())
		}
		addr := crypto.PubkeyToAddress(sk.PublicKey)
		if bytes.Equal(addr.Bytes()[20-numZeros:], bs) {
			fmt.Printf("Private Key: %x\n", crypto.FromECDSA(sk))
			fmt.Printf("Public Key: %x\n", crypto.FromECDSAPub(&sk.PublicKey))
			fmt.Printf("ETH Address: %s\n", addr.Hex())
			done <- nil
		}
	}
}
