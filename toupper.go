package piscine

func ToUpper(s string) string {
	tupper := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			tupper += string(s[i] - ('a' - 'A'))
		} else {
			tupper += string(s[i])
		}
	}
	return tupper
}
