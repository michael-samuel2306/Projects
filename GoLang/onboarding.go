package main
import "fmt"

/* This programme is design by Michael Samuel  to determine (using age 18 - 35 years) 
if an applicant is eligible to be onboarded into the NITDA Blockchain Scholarship.
*/

func Hello(name, state string) string {
	message := fmt.Sprintf("Hi, %v from %v State. Welcome!\n", name, state)
	return message
}

func sorry(name string) string {
	message := fmt.Sprintf("Sorry, %v! You're not eligible for the NITDA Blockchain Scholarship programme.\n", name)
	return message
}

func qualified(name string) string {
	message := fmt.Sprintf("Congratulations, %v! You're eligible for the NITDA Blockchain Scholarship programme.\n", name)
	return message
}


func main(){
	var name, state string
	fmt.Printf("Please enter your first name: \n")
	fmt.Scanf("%s\n", &name)
	fmt.Printf("Please enter your state of origin: \n")
	fmt.Scanf("%s\n", &state)
	fmt.Printf(Hello(name, state))

	var ageNum int
	fmt.Printf("Please enter your age: \n")
	fmt.Scanf("%d", &ageNum)
	if ageNum < 18 || ageNum > 35 {
		fmt.Printf(sorry(name));
	} else {
		fmt.Printf(qualified(name));
		fmt.Printf("Please sign up on 5thwork. Here is the signup link: https://5thwork.com/courses/course-v1:NITDA+NITDA0001+2022Q3/about\n")
		fmt.Printf("Complete the Introduction to Bitcoin Theory course inorder to move to the next phase of this programme. Thank you!\n")
	}
}