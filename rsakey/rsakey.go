package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"math/big"
)

type msgType struct {
	Data json.RawMessage
	Type string
}

func main() {
	fmt.Println("Start!")
	rsakey, errKey := rsa.GenerateKey(rand.Reader, 2048)
	if errKey == nil {
		publicKey := rsakey.Public()               // get the public key from the rsa key
		rsaPublicKey := publicKey.(*rsa.PublicKey) // convert to rsa publick key type
		// fmt.Println("N : ", rsaPublicKey.N)
		// fmt.Println("E : ", rsaPublicKey.E)
		// send N
		msg := map[string]interface{}{"data": rsaPublicKey.N, "type": "publicKey"}
		encoded, errMarshal := json.Marshal(msg) // encode the N param
		if errMarshal == nil {
			var decoded msgType
			errUnmarshall := json.Unmarshal(encoded, &decoded)
			if errUnmarshall == nil && decoded.Type == "publicKey" {
				receivedPublic := big.NewInt(0)
				errUnmarshallkey := json.Unmarshal(decoded.Data, &receivedPublic)
				if errUnmarshallkey == nil {
					fmt.Println(receivedPublic.Cmp(rsaPublicKey.N) == 0) // To check got the correct N or not
				} else {
					fmt.Println("Error : ", errUnmarshallkey)
				}
			} else {
				fmt.Println("UnMarshal error", errUnmarshall)
			}
		} else {
			fmt.Println("Marshal error", errMarshal)
		}
	} else {
		fmt.Println("Rsa key generation Error!", errKey)
	}
}
