// # M 1.3 Largest number
// # Create a function that takes a list of integers as arguments and returns the largest number.
// # You must not use max()

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

func inputNumbers(numbers []int) []int{
	for i:=0; i<10; i++ {
		number := getInputNumber("input number " + (strconv.Itoa(i)) + " : ")
		numbers = append(numbers, number) 
	}

	return numbers
}

func getLargestNumber(numbers []int) int{
	if len(numbers) == 0 { return 0 }
	maxInt := numbers[0]
	for _,number := range(numbers) {
		if (number > maxInt){
			maxInt = number
		}

	}
	return maxInt
}
 
func main(){ 

	numbers := []int{}
	numbers = inputNumbers(numbers)

	fmt.Println("\nThe largest number is :" + strconv.Itoa(getLargestNumber(numbers)))
	

	fmt.Println("\n-- End program Medium1.3 --\n")
}
