package main

import (
	"fmt"
)

func main() {
	 var m map[string][]string
	 fmt.Println(m)

	 m = map[string][]string {
		"coffee": {"Coffee", "Espresso", "Cappuccino"},
		"tea": {"Hot Tea","Chai Tea"},
	 }

	 fmt.Println(m["coffee"])
	 fmt.Println(m["tea"])

	 m["other"] = []string{"Hot Chocolate"}
	 fmt.Println(m) //map[coffee:[Coffee Espresso Cappuccino] other:[Hot Chocolate] tea:[Hot Tea Chai Tea]]

	 delete(m, "tea")
	 v, ok := m["tea"];
	 fmt.Println(v, ok)  //[] false

	 m2 := m  // pointer
	 m2["coffee"] = []string{"Coffee"}
	 m["tea"] = []string{"Hot Tea"}

	 fmt.Println(m) //map[coffee:[Coffee] other:[Hot Chocolate] tea:[Hot Tea]]
	 fmt.Println(m2) //map[coffee:[Coffee] other:[Hot Chocolate] tea:[Hot Tea]]

}
