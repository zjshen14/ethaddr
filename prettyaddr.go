package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

var numZero int

func init()  {
	flag.IntVar(&numZero, "num-zero", 0, "Number of zero bytes at the end of address")
}
func main() {
	chs := make(chan interface{})
	done := make(chan interface{})
	defer close(chs)
	for {
		select {
		case <-done:
			return
		case chs<-nil:
		default:
			go genKey(numZero, chs, done)
		}

	}
}

func genKey(numZero int, chs <-chan interface{}, done chan<- interface{}) {
	bs := make([]byte, numZero)
	for range chs {
		sk, err := crypto.GenerateKey()
		if err != nil {
			log.Panic(err.Error())
		}
		addr := crypto.PubkeyToAddress(sk.PublicKey)
		if bytes.Equal(addr.Bytes()[20-numZero:], bs) {
			fmt.Printf("Private Key: %x\n", crypto.FromECDSA(sk))
			fmt.Printf("Public Key: %x\n", crypto.FromECDSAPub(&sk.PublicKey))
			fmt.Printf("ETH Address: %s\n", addr.Hex())
			done <- nil
		}
	}
}
