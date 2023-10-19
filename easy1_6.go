// # E 1.6 Smallest and largest (Lists)
// # 1.Creates a list of 10 random integers.
// # 2.Looking for a built-in function to find the largest number.
// # 3.Looking for a function to find the smallest number.
// # 4.Finds the sum of all numbers in the list.
// # 5.Sorts the list from lowest to highest.

package main

import (
	"fmt"
	"sort"
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

func getSmallestNumber(numbers []int) int{
	if len(numbers) == 0 { return 0 } 

	minInt := numbers[0]
	for _,number := range(numbers) {
		if (number < minInt){
			minInt = number
		}
	}
	return minInt
}

func sumNumbers(numbers []int) int {
	sum := 0
    for _, number := range numbers {
        sum += number
    }
	return sum
}

func sortMinToMax(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}
 
func main(){
	numbers := []int{}

	fmt.Println("\n1.Creates a list of 10 random integers.")

	for i:=0; i<10; i++ {
		number := getInputNumber("input number " + (strconv.Itoa(i)) + " : ")
		numbers = append(numbers, number) 
	}

	fmt.Println("\nall input numbers")
	for index, number := range numbers {
        fmt.Println(strconv.Itoa(index) + " > " + strconv.Itoa(number))
    }


	fmt.Println("\n2.Looking for a built-in function to find the largest number.")
	fmt.Println("maximum number : " + strconv.Itoa(getLargestNumber(numbers)))

	fmt.Println("\n3.Looking for a function to find the smallest number.")
	fmt.Println("minimum number : " + strconv.Itoa(getSmallestNumber(numbers)))

	fmt.Println("\n4.Finds the sum of all numbers in the list.")
	fmt.Println(strconv.Itoa(sumNumbers(numbers)))

	fmt.Println("\n5.Sorts the list from lowest to highest.")
	for index, number := range sortMinToMax(numbers) {
        fmt.Println(strconv.Itoa(index) + " > " + strconv.Itoa(number))
    }


	fmt.Println("\n-- End program Easy1.6 --\n")
}
