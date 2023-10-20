// N 1.2 Longest valid parenthesis
// Given a string containing just the characters '(' and ')',
// return the length of the longest valid (well-formed) parentheses
// substring.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inputString(text string) string {
	for {
		var input string
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(text)
	
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("input error!! try again")	 
		}

		return strings.TrimSpace(input)
	}
}

func isCorrectFormat(inputString string) bool {
	for _, char := range inputString { 
		srtChar := string(char)
		if srtChar != ")" && srtChar != "(" {
			return false
		}
	}
	return true
}

func longestValidParentheses(inputString string) {
	lengthString := 0
	open := []int{}
	close :=[]int{}

	for index, char := range inputString { 
		srtChar := string(char)
		// Look for '(' and ')'
		if srtChar == "(" {
			open = append(open,index)
		} else if srtChar == ")" {
			close = append(open,index)
		}
	}

	

	if (len(open) == 0 || len(close) == 0) {
		lengthString = 0
	} 

	lastItem := close[len(close)-1] 
	substring := inputString[open[0]:lastItem]
	lengthString = len(substring)
	

	fmt.Println("index of '(' : ")
	fmt.Println(open)
	fmt.Println("index of ')' :")
	fmt.Println(close)

	fmt.Println("\nThe longest valid parentheses substring is '" + substring + "'")
	fmt.Println("longest Valid Parentheses :" + strconv.Itoa(lengthString)+ "\n\n")
	 
}

func main() {

	inputString := inputString("In put the characters only '(' or ')' : ")

	if (isCorrectFormat(inputString)) {
		longestValidParentheses(inputString)
	} else {
		fmt.Println("The string should contain only '(' and ')'")
	}

	fmt.Printf("-- End program Nightmare1.2 --")
}


// Example 1:
// Input: s = "(()"  
// Output: 2
// Explanation: The longest valid parentheses substring is "()".
// Example 2:

// Input: s = ")()())"
// Output: 4
// Explanation: The longest valid parentheses substring is "()()".
// Example 3:

// Input: s = ""
// Output: 0
 
// Constraints:

// 0 <= s.length <= 3 * 104
// s[i] is '(', or ')'.