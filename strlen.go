package piscine

func StrLen(str string) int {
	i := 0
	for index, _ := range str {
		i = index + 1
	}
	return i
}
