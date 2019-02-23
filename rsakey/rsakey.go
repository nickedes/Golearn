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
		// fmt.Println("N : ", rsaPublicKey.N)
		// fmt.Println("E : ", rsaPublicKey.E)
		// send N
		encoded, errMarshal := json.Marshal(rsaPublicKey.N) // encode the N param
		if errMarshal == nil {
			decoded := big.NewInt(0)
			errUnmarshall := json.Unmarshal(encoded, &decoded)
			if errUnmarshall == nil {
				fmt.Println(decoded.Cmp(rsaPublicKey.N) == 0)
			} else {
				fmt.Println("UnMarshal error", errUnmarshall)
			}
		} else {
			fmt.Println("Marshal error", errMarshal)
		}
	} else {
		fmt.Println("Rsa key generation Error!", err)
	}
}
