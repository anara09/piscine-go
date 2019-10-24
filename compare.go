package piscine

func Compare(a, b string) int {
	m := []rune(a)
	l := []rune(b)
	n := 0
	k := 0
	for range m {
		n++
	}
	for range l {
		k++
	}
	for i := 0; i <= n-k; i++ {
		if b == a[i:i+k] && n == k {
			return 0
		} else if b == a[i:i+k] && n > k && b[0] == a[0] {
			return 1
		}
	}
	return -1
}
