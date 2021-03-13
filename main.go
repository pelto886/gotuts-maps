package main

import "fmt"

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
