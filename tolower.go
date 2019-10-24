package piscine

func ToLower(s string) string {
	tlower := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			tlower += string(s[i] + ('a' - 'A'))
		} else {
			tlower += string(s[i])
		}
	}
	return tlower
}
