package main
import "fmt"

func main(){
	
	// Giving a condition using the if-statement. This is a single line comment.


    /* This is a multiple line comment. I don't want this block of code to run.
	y := 10
	if (y > 0) {
		fmt.Printf("Positive\n")
	}
	*/

	/* This is a multiple line comment. I don't want this block of code to run.
	for i := 0; i < 10; i++ {
		// This code will print "Hi" 10x
		fmt.Printf("Hi\n")
	}
	*/

	/* This is a multiple line comment. I don't want this block of code to run.
	i := 0
	for i < 10 {
		// This block of code will print "Hi" for 5x and stop after the fifth time. 
		i++
		if i == 5 {break}
		fmt.Printf("Hi\n")
	}
	*/
	

	/* This is a multiple line comment. I don't want this block of code to run.
	i := 0
	for i < 10 {
		// This block of code will print "Hi" for 5x, skip the fifth one and continue the loop. 
		i++
		if i == 5 {continue}
		fmt.Printf("Hi\n")
	}
	*/
	

	scores := [...]int {10,11,12,13,14,15,16,17,18,19,20}
	
	for i, v := range scores {
		v++
		if v == 15 {continue}
		fmt.Printf("Score %d: %d\n", i, v)
	}

}