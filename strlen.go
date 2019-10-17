package piscine

func StrLen(str string) int {
	i := 0
	for index, _ := range str {
		index++
		i = index
	}
	return i
}
