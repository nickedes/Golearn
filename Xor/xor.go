package main

import (
	"fmt"
	random "math/rand"
	"strconv"
	"time"
)

func sample(A []string, x int) []string {
	random.Seed(time.Now().UnixNano())
	randomize := random.Perm(len(A))

	sampleslice := make([]string, 0)

	for _, v := range randomize[:x] {
		sampleslice = append(sampleslice, A[v])
	}
	return sampleslice
}

func xorbinary(A []string) int64 {
	var xorVal int64
	for i := range A {
		intval, _ := strconv.ParseInt(A[i], 2, 0)
		xorVal = xorVal ^ intval
	}
	return xorVal
}

func getXor(A []string) (string, []string) {
	randomset := sample(A, 2)
	xorVal := xorbinary(randomset)
	xorString := string(xorVal)
	return xorString, randomset
}

func main() {
	fmt.Println(1)
}
