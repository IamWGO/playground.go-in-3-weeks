/*
E 1.1 Age verification (If-else)
1. Write a function that takes in the user's age
2. and then prints if they are minors (under 18),
3. adults (between 18 and 65) or pensioners (over 65).*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getInputString(text string) string {
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

func longUnqueSubstring(inputString string) {
	if len(inputString) == 0 {
		return
	}

	start := 0
	end := 0
	maxStart := 0
	maxEnd := 0
	max_length := 0
	slicesChar := []string{}
 
	for {
		if len(inputString) == 0{ break }
		if end >= len(inputString) { break }

		srtChar := string(inputString[end])

		isContains := slices.Contains(slicesChar, srtChar)

		if isContains == false {
			slicesChar = append(slicesChar, srtChar)

			fmt.Println(slicesChar)

			current_length := end - start
			end++

			if current_length > max_length {
				max_length = current_length
				maxStart = start
				maxEnd = end
			} 
			
		} else {
			slicesChar = slicesChar[1:]		
			start++	
		}
	}

	fmt.Println("\n\n:> maxStart = " + strconv.Itoa(maxStart) + " maxEnd = " + strconv.Itoa(maxEnd))
	fmt.Println(inputString[maxStart:maxEnd])
}

func main() {

	inputString := getInputString("Input some letters (ex: pwwkew)? ")
	longUnqueSubstring(inputString)

	fmt.Printf("-- End program Extream1.1 --")
}
