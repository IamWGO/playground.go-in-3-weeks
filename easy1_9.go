// # E 1.9 The sum of two lists (Lists, Loops)
// # Create a function that takes two equally long lists as arguments.
// # Return a list where each element is the sum of the two lists' respective elements.

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
	numbers1 := []int{}
	numbers2 := []int{}
	numbers3 := []int{}

	fmt.Println("\n1.Creates a list of 10 random integers.")

	fmt.Println("\n Set 1 : ")
	for i:=0; i<5; i++ {
		number := getInputNumber("input number " + (strconv.Itoa(i)) + " : ")
		numbers1 = append(numbers1, number) 
	}

	fmt.Println("\n Set 2 : ")
	for i:=0; i<5; i++ {
		number := getInputNumber("input number " + (strconv.Itoa(i)) + " : ")
		numbers2 = append(numbers2, number) 
	}


	fmt.Println("\n The sum of the two lists : ")
	for i:=0; i<5; i++ {
		numbers3 = append(numbers3, numbers1[i] + numbers2[i])
	}

	for index, sumNumbers := range numbers3 {
		fmt.Println(strconv.Itoa(index) + ". " +
		strconv.Itoa(numbers1[index]) + " + " +
		strconv.Itoa(numbers2[index]) + " = " +
		strconv.Itoa(sumNumbers) )
	}

	fmt.Println("\n-- End program Easy1.9 --\n")
}
