package main

import (
	"fmt"
)

func main() {
	
	if i := 5; i < 5 {
		fmt.Println("i is less than 5")
	} else if i < 10 {
		fmt.Println("i is less than 10")
	} else {
		fmt.Println("i is greater than or equal to 10")
	}
	

}
