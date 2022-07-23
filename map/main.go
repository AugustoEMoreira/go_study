package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff00009",
		"blue":  "#ff00000",
		"white": "#fffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, "is", value)
	}
}
