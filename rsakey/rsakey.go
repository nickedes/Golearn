package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {
	fmt.Println("Start!")
	rsakey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err == nil {
		publicKey := rsakey.Public() // get the public key from the rsa key
		rsaPublicKey := publicKey.(*rsa.PublicKey) // convert to rsa publick key type
		fmt.Println("N : ", rsaPublicKey.N)
		fmt.Println("E : ", rsaPublicKey.E)
	} else {
		fmt.Println("Rsa key generation Error!", err)
	}
}
