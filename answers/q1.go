package answers

func Day1() []interface{} {
	data := ReadInputAsInt(1)
	return []interface{}{q1part1(data), q1part2(data)}

}

func cumulative_sum(data []int) []int {
	cumsum := []int{}
	currentVal := 0
	for _, value := range data {
		if value != 0 {
			currentVal += value
		} else {
			cumsum = append(cumsum, currentVal)
			currentVal = 0
		}
	}
	return cumsum
}

func q1part1(data []int) int {
	cumsum := cumulative_sum(data)
	_, max_val := maxSlice(cumsum)
	return max_val
}

func q1part2(data []int) int {
	cumsum := cumulative_sum(data)
	// Sum top 3
	sum := 0
	for i := 0; i < 3; i++ {
		pos, max_val := maxSlice(cumsum)
		sum += max_val
		cumsum[pos] = 0
	}
	return sum
}
