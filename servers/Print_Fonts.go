package ascii_art

import "fmt"

func Print_Fonts(OutputFonts [][]string) {
	if len(OutputFonts) == 0 {
		return
	}
	for i := 0; i < 8; i++ {

		for _, v := range OutputFonts {
			fmt.Print(v[i])
		}
			fmt.Println()
	}
}
