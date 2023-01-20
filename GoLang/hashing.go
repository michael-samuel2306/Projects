package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	/* This code will produce the hash of the input
	given to the machine */

	var input string
	fmt.Printf("Enter the input you want to hash: \n")
	fmt.Scanf("%s\n", &input)
	hash_result := sha256.Sum256([]byte(input))
	fmt.Printf("Hash: %x", hash_result)
}