package piscine

import "ft"

func PrintStr(s string) {
	i := 0
	for range s {
		ft.PrintRune(rune(s[i]))
		i++
	}
}
