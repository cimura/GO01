package	piscine

import "ft"

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		ft.PrintRune(rune(s[i]))
	}
}