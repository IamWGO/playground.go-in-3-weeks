// # E 1.7 Combine Lists (Lists)
// # Create a function that takes two lists of integers as arguments
// # and returns a combined and sorted list of them.

package main

import (
	"fmt"
	"sort"
)

func combineTwoSlices(numbers1 []int, numbers2 []int) []int {
	combined := []int{}
	
	combined = append(combined, numbers1...)
	combined = append(combined, numbers2...)
    
	sort.Ints(combined)

	return combined
}

 
func main(){

	numbers1 := []int{1, 2, 3}
    numbers2 := []int{4, 10, 6}

    result := combineTwoSlices(numbers1, numbers2)

    fmt.Println(result)

	fmt.Println("\n-- End program Easy1.7 --\n")
}
