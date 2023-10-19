// # M 1.8 Counting with Dictionaries
// # Given a long string of words, create a dictionary that lists the number of each word in the string.
// #

package main

import (
	"bufio"
	"fmt"
	"os"
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

 
func main(){ 
	inputString := inputString("input a string : ")
	wordCounts := map[string]int{}
	// Split the string into words
	words := strings.Fields(inputString)

	// Count word occurrences
	for _, word := range words {
		wordCounts[word]++
	}

	// Print word counts
	for word, count := range wordCounts {
		fmt.Printf("word: %s -> %d\n", word, count)
	}
	fmt.Println("\n-- End program Medium1.8 --\n")
}
