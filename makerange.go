package piscine

func MakeRange(min, max int) []int {
	size := max - min
	if min >= max {
		return nil
	} else {
		ans := make([]int, size)
		for i := 0; i < size; i++ {
			ans[i] = min + i
		}
		return ans
	}

}
