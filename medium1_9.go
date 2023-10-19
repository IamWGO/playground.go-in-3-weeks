// # M 1.9 Invert a Dictionary
// # Create a dictionary of objects where the keys are unique, but the values are not.
// # Write a function to invert the dictionary, by making the values
// # from the original dictionary to keys and the keys from the original dictionary
// # to the host in the new one.


package main

import (
	"fmt"
)
 
func movieTableToTimeTable(movieTable map[string]string) map[string][]string {
	timeTable := map[string][]string{}
	 
	for key,value := range movieTable {
		timeTable[value] = append(timeTable[value], key)
		// if _, exists := timeTable[value]; exists {
		// }
	}
	
	return timeTable
}

func printMovieTable(movieTable map[string]string){ 
	for key,value := range movieTable {
		fmt.Println("  Title : " + key  + "  show time : " + value)
	}
}

func printTimeTable(timeTable map[string][]string) {
	for key,movies := range timeTable {
		fmt.Println("  Show time : " + key )
		for _,movie := range movies {
			fmt.Println("     " + movie )
		}
	}
}

 
func main(){ 

	movieTable := map[string]string{}
	movieTable["movie1"] = "11:00am"
	movieTable["movie2"] = "11:00am"
	movieTable["movie3"] = "3:00pm"
	movieTable["movie4"] = "5:00pm"

	fmt.Println("\nAll movies")
	printMovieTable(movieTable)

	fmt.Println("\nTime table")
	//map[string][]string{}
	timeTable := movieTableToTimeTable(movieTable)
	printTimeTable(timeTable)


	fmt.Println("\n-- End program Medium1.1 --\n")
}
