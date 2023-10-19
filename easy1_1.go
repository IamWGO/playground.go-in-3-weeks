/*
E 1.1 Age verification (If-else)
1. Write a function that takes in the user's age
2. and then prints if they are minors (under 18),
3. adults (between 18 and 65) or pensioners (over 65).*/

package main

import (
	"fmt"
	"strconv"
)

func getInputNumber(text string) int {
	for {
		fmt.Print(text)
		input := ""
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Empty!! input some number")
			continue
		}
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("The input is not a valid integer.")
			continue
		}
		return number
	}
}

func main() {

	var age = getInputNumber("How old are you? ")
	fmt.Printf("Your age is: %v ", age)
 
	if age < 18 {
	fmt.Println("minderåriga")
	} else if age >= 18 && age <= 65 {
		fmt.Println("mellan 18 och 65")
	} else {
		fmt.Println("över 65")
	}

	fmt.Printf("-- End program Easy1.1 --")
}
