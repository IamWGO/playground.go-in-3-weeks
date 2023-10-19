// # M 1.5 FizzBuzz
// # FizzBuzz is a classic interview problem.
// #
// # Write a function that prints all numbers from 1 to a given number.
// # If the number is divisible by 3, print "Fizz". If the number is divisible by 5, print "Buzz".
// # If the number is divisible by 3 and 5, print "FizzBuzz".

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

func printNumbers(maxLoop int) {
	for i:=1; i<=maxLoop; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(strconv.Itoa(i) + ". FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(strconv.Itoa(i) + ". Fizz")
		} else if i%5 == 0 {
			fmt.Println(strconv.Itoa(i) + ". Buzz")
		} else {
			fmt.Println(strconv.Itoa(i) + ". -")
		}
	}
}
 
func main(){ 

	inputNumbers := getInputNumber("Which number do you want to list (1..n)? ")

	printNumbers(inputNumbers)

	fmt.Println("\n-- End program Medium1.5 --\n")
}
