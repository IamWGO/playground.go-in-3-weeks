// # E 1.3 - Interval (If-else)
// 1 Create the function in_range(lower, upper)
// 2 which determines if a number is within the specified range.
// 3 If the number is in the endpoints, it is counted as within the range.


package main

import (
	"fmt"
	"strconv"
)

func getInputNumber1(text string) int {
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

func isNumberInRange(number int) {
	stringNumber := strconv.Itoa(number)
	if number>=10 || number<=20 {
		fmt.Printf("number %s is in range.\n", stringNumber)
	} else {
		fmt.Printf("number %s is not in range 10 and 20.\n", stringNumber)
	}
}

func main(){
	var inputNumber = getInputNumber1("Your number? ")
	isNumberInRange(inputNumber)

	fmt.Printf("-- End program Easy1.3 --")
}
