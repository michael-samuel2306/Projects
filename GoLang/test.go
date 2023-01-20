package main

import (
	"fmt"
	"crypto/sha256"
)

func proof_of_work(header, difficulty_bit int) int {
	a := (256 - difficulty_bit)
	difficulty_target := 2 ** a
	return difficulty_target
}

func main(){

		max_nonce := 2 ** 32
	
		for nonce := range max_nonce {
			hashResult := sha256.Sum256([]byte("strconv.Itoa(header)"))
			fmt.Printf("%x", hashResult)
	
			// data := []byte("nJ1m4Cc3")
			// fmt.Printf("%x", sha1.Sum(data))
	
	
			if hashResult <= difficulty_target {
				fmt.Printf("Success with nonce %d", nonce)
				fmt.Printf("Hash is %s",  hashResult)
				// return (hashResult, nonce)
			}
			fmt.Printf("Failed after %d {max_nonce} tries", nonce)
			return nonce
		}
}