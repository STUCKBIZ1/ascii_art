package main

import (
	"fmt"
	"os"

	ascii_art "ascii_art/src"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: input txt")
		os.Exit(1)
	}
	inputdata := os.Args[1]
	for _, char := range inputdata {
		if char < ' ' || char > '~' {
			fmt.Println("error: this unsopported rune")
			os.Exit(1)
		}
	}
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	res1 := ascii_art.Sep_Fonts(string(data))
	ascii_art.Chars_To_Art(res1, inputdata)
}
