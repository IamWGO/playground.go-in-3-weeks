// # E 1.10 Basic key-value pairing
// 1. Create a dictionary that represents a person;
// 2. it must contain first name, last name, age and email address.
// 3. Print each piece of information individually using the keys.
// 4. Add a new key-value pair representing the person's hometown.
// 5. Check if the person has a key "middlename". If not, add it with a value of your choice.
// 6. Update the person's age by incrementing it by one.

package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Person struct {
    firstName string
	middlename string
    lastName  string
    age       int
	hometown string
    email     string
}

func getInputString(text string) string {
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

func printInformation(person Person){
	val := reflect.ValueOf(person)
	typ := reflect.TypeOf(person)

	for i := 0; i < val.NumField(); i++ {
		 fieldName := typ.Field(i).Name
		 fieldValue := val.Field(i)
		 
		if val.Field(i).Kind() == reflect.Int {
			show := strconv.Itoa(int(fieldValue.Int())) 
			fmt.Println(fieldName + " -> " + show)
		} else {
			fmt.Println(fieldName + " -> " + val.Field(i).String())
		} 
	}
}
 
func main(){
	fmt.Println("\n1. Create a dictionary that represents a person" +
				"\n2. it must contain first name, last name, age and email address.")
	person := Person{
		firstName: "Waleeat",
		lastName: "Gottlieb",
		age: 44,
		email: "waleeat.gottlieb@gmail.com",
	}

	fmt.Println(person)

	fmt.Println("\n3. Print each piece of information individually using the keys.")
	printInformation(person)

	fmt.Println("\n4. Add a new key-value pair representing the person's hometown.")
	hometown := getInputString("In put your hometown: ")
	person.hometown = hometown
	printInformation(person)
	
	fmt.Println("\n5. Check if the person has a key \"middlename\". If not, add it with a value of your choice.")
	answer := getInputString("Do you have middle name? (y/any): ")
	if strings.ToLower(answer) == "y" {
		hometown := getInputString("In put your middle name: ")
		person.middlename = hometown
	}
	printInformation(person)

	fmt.Println("\n6. Update the person's age by incrementing it by one.")
	person.age += 1
	printInformation(person)

	fmt.Println("\n-- End program Easy1.10 --\n")
}
