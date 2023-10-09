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
}
