package	piscine

func Strlen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func Atoi(s string) int {
	num := 0
	i := 0
	sign := 1
	flag := 0
	if rune(s[i]) == '+' {
		i++
		flag++
	} else if rune(s[i]) == '-' {
		sign = -1
		flag++
		i++
	}
	for i < Strlen(s)-1 + flag {
		check := int(s[i] - '0')
		if check >= 0 && check <= 9 {
			num = 10 * num + check
			i++
		} else {
			return 0
		}
	}
	return num * sign
}