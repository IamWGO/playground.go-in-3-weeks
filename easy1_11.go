// # E 1.11 Loop through Dictionaries (Dictionaries, Loops)
// # 1. Create a dictionary where the keys are numbers between 1 and 15 and the values are the squares of the keys.
// # 2. Print each key-value pair on a new line.
// # 3. Calculate and print the sum of all values in the dictionary.

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

func sumMapNumbers(mapNumbers map[int]int) int{
	sum := 0
	for _,value := range mapNumbers {
		sum += value
	}

	return sum
}

func printMapNumbers(mapNumbers map[int]int){ 
	for key,value := range mapNumbers {
		fmt.Println(strconv.Itoa(key)  + ":" + strconv.Itoa(value))
	}
}

 
func main(){

	fmt.Println("\n1. Create a dictionary where the keys are numbers between 1 and 15 and the values are the squares of the keys.")
	// Create an empty map with string keys and int values
	mapNumbers := make(map[int]int)

	for i:=1; i<=15; i++ {
		number := getInputNumber("input number " + (strconv.Itoa(i)) + " : ")
		mapNumbers[i] = number
	}

	fmt.Println("\n2. Print each key-value pair on a new line.")
	printMapNumbers(mapNumbers)

	fmt.Println("\n3. Calculate and print the sum of all values in the dictionary.")
	fmt.Println("Total : " + (strconv.Itoa(sumMapNumbers(mapNumbers))))

	 





	fmt.Println("\n-- End program Easy1.11 --\n")
}
