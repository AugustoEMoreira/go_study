package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":  "#ff0000",
	// 	"blue": "#ff0000",
	// }
	// var colors map[string]string
	colors := make(map[string]string)
	colors["white"] = "#fffff"

	delete(colors, "white")
	fmt.Println(colors)
}
