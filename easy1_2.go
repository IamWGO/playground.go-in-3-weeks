/*
 E 1.2 Even or odd (If-else)
1. Write a function that takes in a single number
2. and prints whether it is odd or even.*/

package main

import (
	"fmt"
	"strconv"
)

func getInputNumber1(text string) int {
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

func oddOrEven(number int) {
	if number%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

func main(){
	var inputNumber = getInputNumber1("Your number?? ")
	fmt.Printf("number is : %v ", inputNumber)
	oddOrEven(inputNumber)
	fmt.Printf("-- End program Easy1.2 --")
}
