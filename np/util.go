package np

func double(f func(y int, x int)) {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			f(y, x)
		}
	}
}

func search(slice []int, n int) (bool, []int) {
	var indexes []int
	var contain = false
	for i, v := range slice {
		if v == n {
			indexes = append(indexes, i)
			contain = true
		}
	}
	return contain, indexes
}

func sliceMerge(slices ...[]int) []int {
	var merged []int
	for _, slice := range slices {
		merged = append(merged, slice...)
	}
	return merged
}

func set(slice []int) []int {
	m := make(map[int]bool)
	var uniq []int
	for _, ele := range slice {
		if !m[ele] {
			m[ele] = true
			uniq = append(uniq, ele)
		}
	}
	return uniq
}
