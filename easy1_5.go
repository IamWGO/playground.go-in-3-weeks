// # E 1.5 List slicing (Lists)
// # Write a function like:
// #
// # 1.Creates a list containing the numbers 1 to 10.
// # 2.Uses list slicing to print the first five numbers.
// # 3.Uses list slicing to print the last four numbers.
// # 4.Prints every other number in the list, starting with the first one.
// # 5.Creates a new list that is a copy of the reversed original list.

package main

import (
	"fmt"
	"strconv"
)

 
func main(){
	fmt.Println("\n1.Creates a list containing the numbers 1 to 10.")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(numbers)

	fmt.Println("\n2.Uses list slicing to print the first five numbers.")
	firstFiveNumbers := numbers[:5]
	
	fmt.Println("\n3.Uses list slicing to print the last four numbers.")
	lastFourNumbers := numbers[6:]

	fmt.Println("\n4.Prints every other number in the list, starting with the first one.")
	fmt.Println("all numbers")
	for index, number := range numbers {
        fmt.Println(strconv.Itoa(index) + " > " + strconv.Itoa(number))
    }
	fmt.Println("first 5 numbers")
	for index, number := range firstFiveNumbers {
        fmt.Println(strconv.Itoa(index) + " > " + strconv.Itoa(number))
    }
	fmt.Println("last 4 numbers")
	for index, number := range lastFourNumbers {
        fmt.Println(strconv.Itoa(index) + " > " + strconv.Itoa(number))
    }

	fmt.Println("\n5.Creates a new list that is a copy of the reversed original list.")
	copyNumbers := numbers
 
	reversedNumbers := []int{}
	for itemNumber:=len(copyNumbers); itemNumber>0; itemNumber-- {
		reversedNumbers = append(reversedNumbers, copyNumbers[itemNumber-1])	
	}
    
	for index, number := range reversedNumbers {
        fmt.Println(strconv.Itoa(index) + " > " + strconv.Itoa(number))
    }

	fmt.Println("\n-- End program Easy1.5 --\n")
}
