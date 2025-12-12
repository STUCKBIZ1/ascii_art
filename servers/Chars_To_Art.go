package ascii_art

import (
	"fmt"
	"strings"
)

func Chars_To_Art(fonts [][]string, input string) {
	var CInput [][]string
	LnIptData := strings.Split(input, "\\n")
	if LnIptData[0] == "" {
		LnIptData = LnIptData[1:]
	}
	for _, inputline := range LnIptData {
		if inputline == "" {
			fmt.Println()
			continue
		}
		for _, char := range inputline {
			for j, font := range fonts {
				if j+32 == int(char) {
					CInput = append(CInput, font)
				}
			}
		}
		Print_Fonts(CInput)
		CInput = [][]string{}
	}
}
