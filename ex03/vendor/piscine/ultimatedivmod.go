package piscine

func UltimateDivMod(a *int, b *int) {
	tmp := *a % *b
	*a = *a / *b
	*b = tmp
}