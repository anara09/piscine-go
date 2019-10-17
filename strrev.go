package piscine

func StrRev(s string) string {
	var s2 string
	for i := range s {
		s2 = string(s[i]) + s2
	}
	return s2
}
