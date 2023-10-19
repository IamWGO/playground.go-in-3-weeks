// # 1.Create a nested list where each element is a list containing three numbers.
// # 2.Prints the first number from each nested list using a loop.
// # 3.Adds a fourth number to each nested list.
// # 4.Uses a loop to print the sum of the numbers in each nested list.

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

 
func main(){ 

	nestedList := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("\n# Prints the first number from each nested list using a loop.")
	fmt.Println("# Adds a fourth number to each nested list.\n")

	for i, nested := range nestedList {
		fmt.Printf("First number of %v is %d\n", nested, nested[0])
		inputNumber := getInputNumber(fmt.Sprintf("%d - Input one more number: ", i))
		nestedList[i] = append(nestedList[i], inputNumber)
	}

	fmt.Println("\n# Uses a loop to print the sum of the numbers in each nested list.")
	for _, nested := range nestedList {
		sum := 0
		for _, num := range nested {
			sum += num
		}
		fmt.Printf("Sum of %v = %d\n", nested, sum)
	}

	fmt.Println("\n-- End program Medium1.1 --\n")
}

