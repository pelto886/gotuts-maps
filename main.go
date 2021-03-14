package main

import "fmt"

// Maps and structs differ in some ways:
// Maps have all keys of same type
// maps have all values of same type
// maps keys are indexed we can iterate
// maps represent a collection of similar properties
// keys can be added after compile time
// maps are a reference type (no pointer needed)

// structs can have values of different type
// structs keys does not support indexing (no iteration)
// structs need to know all the different fields at compile
// used to represent an entity with different properties
// structs are value types, need pointers.
func main() {
	var testColors map[string]string

	testColorsTwo := make(map[string]string)

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	testColorsTwo["white"] = "#FFFFFF"

	fmt.Println(colors)
	fmt.Println(testColors)
	fmt.Println(testColorsTwo)

	delete(testColorsTwo, "white")

	fmt.Println(testColorsTwo)

	printMapContent(colors)

}

func printMapContent(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
