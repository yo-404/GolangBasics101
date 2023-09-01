package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	// "math/rand"
)

func main() {
	// using math/rand
	// fmt.Println("Random number generated is ", rand.Intn(6))

	// generating numbers through cryptography - crpto/rand
	randomnum, err := rand.Int(rand.Reader, big.NewInt(200))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("random number generated from crypto package is :", randomnum)

}
