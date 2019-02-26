package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	printMap(colors)
}
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

//**** Map **** //
// all keys must be the same type
// all values must be the same type
// all keys are indexed, we can iterate over all of them
// REFERENCE TYPE
// use case : when you are representing a very closely related values
//dont need to define all property names at compile time
// ^^ we can add and remove keys as we please
//
//

//**** Struct **** //
// values can be of different type
// keys dont support indexing
// we can not iterate of them
// VALUE TYPE
// we need to strictly define each key before compile time and the value type
// use case
