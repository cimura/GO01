package	piscine

func BasicAtoi(s string) int {
	num := 0
	i := 0
	for range s {
		num = 10 * num + (int(s[i]) - '0')
		i++
	}
	return num
}