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

type IDENTITY struct {
	IP              string
	PK              rsa.PublicKey
	CommitteeID     int64
	PoW             map[string]interface{}
	EpochRandomness string
	Port            int
}

type Dmsg struct {
	Identity IDENTITY
}

func test(rsakey *rsa.PrivateKey) {
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
}

func main() {
	fmt.Println("Start!")
	rsakey, errKey := rsa.GenerateKey(rand.Reader, 2048)
	if errKey == nil {
		test(rsakey)
	} else {
		fmt.Println("Rsa key generation Error!", errKey)
	}

	PK := rsakey.PublicKey

	a := make([]int, 0)
	PoW := map[string]interface{}{"nonce": 12, "hash": "31718973312319203130abcdefd", "list": a}

	identity := IDENTITY{IP: "2.0.1.0", PK: PK, CommitteeID: 2, PoW: PoW, EpochRandomness: "0011", Port: 49125}

	data1 := Dmsg{identity}
	mp := map[string]interface{}{"data": data1, "type": "lol"}

	data, err := json.Marshal(mp)
	if err == nil {
		var id msgType
		_ = json.Unmarshal(data, &id)

		var identity2 Dmsg

		_ = json.Unmarshal(id.Data, &identity2)
		fmt.Println(identity2.Identity)
	}

}
