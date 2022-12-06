package answers

func Day6() []interface{} {
	data := ReadInputAsStr(6)
	return []interface{}{q6part1(data[0]), q6part2(data[0])}
}

func checkUniqueness(data string, window int) int {
	for i := window; i < len(data); i++ {
		bailout := false
		for j := i - window; j <= i; j++ {
			for k := i - window; k < j; k++ {
				if j == k {
					continue
				}
				if data[j] == data[k] {
					bailout = true
					break
				}
			}
			if bailout == true {
				break
			}
		}

		if bailout == true {
			continue
		}
		return i + 1
	}
	return 0
}

func q6part1(data string) int {
	return checkUniqueness(data, 3)
}

func q6part2(data string) int {
	return checkUniqueness(data, 13)
}
