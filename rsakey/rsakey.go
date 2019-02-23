package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Start!")
	rsakey, errKey := rsa.GenerateKey(rand.Reader, 2048)
	if errKey == nil {
		publicKey := rsakey.Public()               // get the public key from the rsa key
		rsaPublicKey := publicKey.(*rsa.PublicKey) // convert to rsa publick key type
		fmt.Println("N : ", rsaPublicKey.N)
		fmt.Println("E : ", rsaPublicKey.E)
	} else {
		fmt.Println("Rsa key generation Error!", err)
	}
}
