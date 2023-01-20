package main

import "fmt"

func main() {
		// This code will calculate your age in the year 2023.
		var birth int
		fmt.Printf("Let's know your age.\n")
		fmt.Printf("Please enter your year of birth? \n")
		fmt.Scanf("%d", &birth)
		// fmt.Printf("%d\n", birth) i don't this line of code to run.
		age := 2023 - birth
		fmt.Printf("You are %d years old.\n", age)
		if age >= 18 {
			fmt.Printf("You are an adult.\n")
		} else {
			fmt.Printf("You are still a child.\n")
		}
}  