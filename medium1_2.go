// # M 1.2 Most frequent number
// # Create a function that takes list as argument,
// # and returns the value that occurs the most times in the list.
// #

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

func numberExists(numbers []int, target int) bool {
    for _, number := range numbers {
        if number == target {
            return true
        }
    }
    return false
}

func  valueOccursMostTimes(numbers []int) map[int]int{
	// Check if the key exists in the map
	occurValues := map[int]int{}
	for _,number := range numbers {
		if numberExists(numbers, number) {
			occurValues[number] += 1
		} else {
			occurValues[number] = 1
		}
	}

	return occurValues
}

func printOccurValues(occurValues map[int]int){ 
	for number,times := range occurValues {
		fmt.Println("  Number : " + strconv.Itoa(number)  + "  Occur : " + strconv.Itoa(times))
	}
}

 
func main(){ 

	numbers := []int{}
	numbers = inputNumbers(numbers)

	occurValues := valueOccursMostTimes(numbers)

	printOccurValues(occurValues)

	fmt.Println("\n-- End program Medium1.2 --\n")
}
