package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("What do you want to say : ")
	in := bufio.NewReader(os.Stdin) // define input 
	s, _ := in.ReadString('\n') // input string and catch error
	s = strings.TrimSpace(s) //trim
	s = strings.ToUpper(s) // set uppercase
	
	fmt.Println(s + "!")
}
