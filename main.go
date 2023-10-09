package main

import "fmt"

func main() {

i := 10
if i<15 {
	fmt.Println(i)
	goto myLabel
}

myLabel : 
	j := 42
	for i := 0; i < j; i++ {
		fmt.Println(i)
	}
 
} 