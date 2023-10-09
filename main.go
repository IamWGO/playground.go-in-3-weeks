package main

import (
	"fmt"
)

func main() {
	var arr [3]int
	fmt.Println(arr)
	arr = [3]int{1,2,3}

	fmt.Println(arr[1])
	arr[1] = 99
	fmt.Println(arr)

	fmt.Println(len(arr));

	fmt.Println("---- String array ----");
	arrString := [3]string{"Lee", "Mali", "Singto"}

	arrString2 := arrString
	fmt.Println(arrString2)

	arrString[0] = "King"
	fmt.Println(arrString)
	fmt.Println(arrString2)
	
	fmt.Println(arrString == arrString2)
	

}
