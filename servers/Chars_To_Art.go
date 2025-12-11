package ascii_art

import (
	"fmt"
	"strings"
)

func Chars_To_Art(fonts [][]string, input string) {
	var CInput [][]string
	LnIptData := strings.Split(input, "\\n")
	for i, inputline := range LnIptData {
		for _, char := range inputline {
			for j, font := range fonts {
				if j+32 == int(char) {
					CInput = append(CInput, font)
				}
			}
		}
		Print_Fonts(CInput)
		if i < len(LnIptData)-1 {
			fmt.Println()
		}
	}
}
