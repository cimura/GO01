package piscine

// import "fmt"

func Strlen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func StrRev(s string) string {
	rev := []rune(s)
	obv := []rune(s)

	count := Strlen(s) - 1
	i := 0
	for range s {
		rev[i] = obv[count]
		count--
		i++
	}
	return string(rev)
}
