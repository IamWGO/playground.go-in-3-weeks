package main

import (
	"fmt"
	"slices" 
)

func main() {
	var s []int  //Slices of ints
	fmt.Println(s) // [] nil
	
	s = []int{1,2,3}
	s[1] = 99
	fmt.Println(s)

	s = append(s, 4,5,6,7)
	fmt.Println(s)

	s = slices.Delete(s, 2,4) // remove index
	fmt.Println(s)

}
