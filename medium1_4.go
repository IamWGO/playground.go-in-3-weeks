// # M 1.4 List backwards
// # Create a function that takes in two lists.
// # The function returns true/false if one list is the other backwards.

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
	for i:=0; i<5; i++ {
		number := getInputNumber("input number " + (strconv.Itoa(i)) + " : ")
		numbers = append(numbers, number) 
	}

	return numbers
}

func ifOtherSliceBackwards(numbers1 []int, numbers2 []int) bool {
	if len(numbers1) != len(numbers1) {
        return false
    }

	for itemNumber := len(numbers1); itemNumber>0; itemNumber-- {
        if numbers1[len(numbers1) - itemNumber] != numbers2[itemNumber-1] {
            return false
        }
    }
	return true
}

 
func main(){ 
	fmt.Println("\nSet 1 :")
	numbers1 := []int{}
	numbers1 = inputNumbers(numbers1)

	fmt.Println("\nSet 2 :")
	numbers2 := []int{}
	numbers2 = inputNumbers(numbers2)

	isOtherSliceBackwards := ifOtherSliceBackwards(numbers1, numbers2)
	if isOtherSliceBackwards {
		fmt.Println("\n One list is the other backwards")
	} else {
		fmt.Println("\n One list is not the other backwards -_-'")
	}

	fmt.Println("\n-- End program Medium1.4 --\n")
}
