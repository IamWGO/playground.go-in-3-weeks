// # E 1.4 - Favorite fruits (Lists)
// # Write a function like:
// #
// # 1.Creates a list containing five elements, which are your favorite fruits.
// # 2. Adds a new fruit to the list.
// # 3.Changes the third element to another fruit.
// # 4.Removes the last element in the list.
// # 5. Prints each fruit in the list on a new line.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"slices"
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
 
func addStringToList(mySlice []string) []string {
	for i := 0; i < 5; i++ {
        input := inputString("input String " + strconv.Itoa(i+1) + " : ")
		mySlice = append(mySlice, input)
    }
	return mySlice
}

func updateStringInList(mySlice []string) []string {
	updateItem := 3
	inputString := inputString("input String : ")

	if len(mySlice) >= updateItem {
		mySlice[updateItem-1] = inputString
	}
	return mySlice
}

func deleteLastItemInList(mySlice []string) []string{
	// Delete from index to items
	//mySlice = slices.Delete(mySlice, 0,3)
	// Delete the last item
    if len(mySlice) > 0 {
        mySlice = mySlice[:len(mySlice)-1]
    }
	return mySlice
}

func printList(mySlice []string){
	for i:=0; i<=len(mySlice)-1; i++ {  
		fmt.Println(strconv.Itoa(i+1) + " > " + mySlice[i]) // Output
	}
}
 
func main(){
	mySlice := []string{}

	fmt.Println("\n=== Input 3 fruits ===")
	mySlice = addStringToList(mySlice)

	fmt.Println("\n=== Changes the third element to another fruit. ===")
	mySlice = updateStringInList(mySlice)
	fmt.Println(mySlice) // Output

	fmt.Println("\n=== Removes the last element in the list. ===")
	mySlice = deleteLastItemInList(mySlice)
	fmt.Println(mySlice) // Output

	fmt.Println("\n=== Prints each fruit in the list on a new line. ===")
	printList(mySlice)

	fmt.Println("\n-- End program Easy1.4 --\n")
}
