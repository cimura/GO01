package	main

import (
	"ft"
	// "fmt"
)

func printArray(n int, first int, end int) {
	for i := 0; i < 
}

func combination(n int, first int, end int) {
	if first == 9 {
		return
	}
	for i := first; i < end + n; i++ {
		ft.PrintRune(rune(i + '0'))
	}
	ft.PrintRune(',')
	ft.PrintRune(' ')
	if (first == 9) {
		first = 0
		combination(n, indx + 1, first)
		fmt.Printf("indx-> %d\n", indx)
	}
	combination(n, indx, first + 1)
}

func PrintCombN(n int) {
	combination(n, 0, 0)
}

func	main() {
	PrintCombN(3)
}