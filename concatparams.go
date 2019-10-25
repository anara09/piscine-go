package piscine

func ConcatParams(args []string) string {
	ans := ""
	for i, c := range args {
		if i == 0 {
			ans = c
			continue
		}
		ans = ans + "\n" + c
	}
	return ans
}
