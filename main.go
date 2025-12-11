package main

import (
	"fmt"
	"os"

	"ascii_art/servers"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: input txt")
		return
	}
	inputdata := os.Args[1]
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("error: ", err)
	}
	res1 := ascii_art.Sep_Fonts(string(data[1:]))
	ascii_art.Chars_To_Art(res1, inputdata)
}
