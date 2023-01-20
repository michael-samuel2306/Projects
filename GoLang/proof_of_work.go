package main

import (
	"fmt"
	"crypto/sha256"
	"time"
	"strconv"
)


var header, nonce, difficulty_bit int

func proof_of_work(header, difficulty_bit int) int {
	difficulty_target := 2 ** (256 - difficulty_bit)
	max_nonce := 2 ** 32

	for nonce := range max_nonce {
		//hashResult := sha1.Sum((str(header) + str(nonce)).encode()).hexdigest()
		hashResult := sha256.New()
		hashResult.Write([]byte(strconv.Itoa(header) + strconv.Itoa(nonce)))
		// fmt.Printf("%x", h.Sum(nil))

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

func main() {

	if __name__ == "__main__" {
    	nonce := 0
    	hashResult := ""

    	for difficulty_bit < 32 {
        	difficulty := 2 ** difficulty_bit;
        	fmt.Printf("Difficulty: %ld (%d bits)", difficulty, difficulty_bit)
        	fmt.Printf("Start searching...")

        	start_time := time.Now().Second()

        	new_block := "test block with transaction" + hashResult

		// find a valid nonce for the block
			hash_result, nonce := proof_of_work(new_block, difficulty_bit)

        	end_time := time.Now().Second()

        	elasped_time := (end_time - start_time)
        	fmt.Printf("Elasped Time: %.4f seconds", elasped_time)

        	if (elasped_time > 0) {
            	hashPower := nonce/elasped_time
            	fmt.Printf("Hashing Power: %ld hashes per second", hashPower)
			}
		}
	}
}