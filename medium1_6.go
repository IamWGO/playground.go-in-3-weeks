// # M 1.6 Every other letter
// # Write a function that takes a string as an argument and returns every other letter.

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

func getEveryOtherLetter(inputString string) string {
	runes := []rune(inputString)
	str := ""
    for index, char := range runes {
		if ((index +1)%2 == 0) {
			str += string(char)
		}
    } 
	return str
}
 
func main(){ 
	inputString := inputString("input a string : ")  
	str := getEveryOtherLetter(inputString)
	fmt.Println(str)
	fmt.Println("\n-- End program Medium1.6 --\n")
}