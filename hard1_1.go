// # Given an array of integers nums and an integer target,
// # return indices of the two numbers such that they add up to target.
// #
// # You may assume that each input would have exactly one solution,
// # and you may not use the same element twice.
// # You can return the answer in any order.
// # Example 1:
// # Input: nums = [2,7,11,15], target = 9
// # Output: [0,1]

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

func inputNumbers(maxLoop int, numbers []int) []int{
	for i:=0; i<maxLoop; i++ {
		number := getInputNumber("input number " + (strconv.Itoa(i)) + " : ")
		numbers = append(numbers, number) 
	}

	return numbers
}
 
func main(){ 
	numbers := []int{}
	inputNumber := getInputNumber("How many numbers do you want to input? ")

	numbers = inputNumbers(inputNumber, numbers)

	targetNumber := getInputNumber("input target:")

	result := [][]int{}

	for i, arg1 := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			sum := arg1 + numbers[j]
			answer := []int{i, j, sum}
			result = append(result, answer)
		}
	}


	targetIndex := [][]int{}

	for _, nested := range result {
		// Check if answer(c) == target_number
		if targetNumber == nested[2] {
			targetList := []int{nested[0], nested[1]}
			targetIndex = append(targetIndex, targetList)
		}
	}

	fmt.Println(targetIndex)



	fmt.Println("\n-- End program Hard1.1 --\n")
}
