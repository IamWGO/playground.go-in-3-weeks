// # E 1.8 Even numbers in lists (Lists, Loops)
// # Create a function that takes in a list of integer as arguments,
// # and returns how many even numbers there are in the list.

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

func getEvenNumberFromSlice(numbers []int) []int{
	return_numbers := []int{}
	//for index, number := range lastFourNumbers {
	for _, number := range numbers {
		if number%2 == 0 {
			return_numbers = append(return_numbers, number)
		}
	} 

	return return_numbers
}

 
func main(){
	numbers := []int{}

	fmt.Println("\n1.Creates a list of 10 random integers.")

	for i:=0; i<10; i++ {
		number := getInputNumber("input number " + (strconv.Itoa(i)) + " : ")
		numbers = append(numbers, number) 
	}

	fmt.Println("\nThe even numbers in the list.")
	for index, number := range getEvenNumberFromSlice(numbers) {
		fmt.Println(strconv.Itoa(index) + " > " + strconv.Itoa(number))
	}

	fmt.Println("\n-- End program Easy1.8 --\n")
}
