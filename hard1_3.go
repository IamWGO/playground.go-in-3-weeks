// # H 1.3 Happy numbers
// # https://leetcode.com/problems/happy-number/

// # Input: n = 19
// # Output: true
// # Explanation:
// # 12 + 92 = 82
// # 82 + 22 = 68
// # 62 + 82 = 100
// # 12 + 02 + 02 = 1
// # Example 2:
// #
// # Input: n = 2
// # Output: false

package main

import (
	"fmt"
	"math"
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

func processHappyNumber(happyNumber int){
	happyNumberString := strconv.Itoa(happyNumber)
	var answer int 
	
	for {
		if len(happyNumberString) == 1 {
			break
		}

		answer = 0
		printRow :=""
		for _, char := range happyNumberString { 
			srtChar := string(char)

			num, _ := strconv.ParseFloat(srtChar, 64)
			 
			answer += int(math.Pow(num, 2))
			printRow += " " + srtChar + "**2" + " +"
		}

		printRow = printRow[:len(printRow)-1] + " = " + strconv.Itoa(answer)
		happyNumberString = strconv.Itoa(answer)

		fmt.Println(printRow)
		 
	}
}
 
func main(){ 
	happyNumber := getInputNumber("Input your happy number ? ")
	

	if len(strconv.Itoa(happyNumber)) > 1 {
		fmt.Println("\nit is happy number ^^")
		processHappyNumber(happyNumber)
	} else {
		fmt.Println("\nit is not happy number!!")
	} 

	
	fmt.Println("\n-- End program Hard1.3 --\n")
}
