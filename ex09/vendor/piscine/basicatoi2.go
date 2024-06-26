package	piscine

func BasicAtoi2(s string) int {
	num := 0
	i := 0
	for range s {
		check := int(s[i] - '0')
		if check >= 0 && check <= 9 {
			num = 10 * num + check
			i++
		} else {
			return 0
		}
	}
	return num
}