package main

import "fmt"

func main() {
	// approach 1
	colors := map[string]string{
		"red":    "#ff0000",
		"green":  "#4bf745",
		"purple": "#38sf62",
	}

	// approach 2
	// var colors map[string]string

	// approach 3
	// colors := make(map[string]string)

	// add key-value pair to an existing map
	// colors["white"] = "#ffffff"

	// delete key-value pair to an existing map
	// delete(colors, "white")
	// fmt.Println(colors)

	printMap(colors)
}

// iterating over maps
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
