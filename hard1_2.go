// # https://leetcode.com/problems/roman-to-integer/

// # Logic form ChatGPT -> refactor roman_dict from array to dictionary

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

func intToRoman(num int) string {
	romanDict := map[int]string{
		1000: "M",
		900:   "CM",
		500:   "D",
		400:   "CD",
		100:   "C",
		90:    "XC",
		50:    "L",
		40:    "XL",
		10:    "X",
		9:     "IX",
		5:     "V",
		4:     "IV",
		1:     "I",
	}

	romanNumeral := ""
	for key, value := range romanDict {
		// Check only if result > 0 (integer division)
		for num >= key {
			// Add the Roman digit
			romanNumeral += value
			// Update the value of num
			num -= key
		}
	}

	return romanNumeral
}

func main() {
	integerValue := getInputNumber("Input a number : ")
	fmt.Printf("Roman numeral: %s\n", intToRoman(integerValue))
 
	fmt.Println("\n-- End program Hard1.2 --\n")
}
