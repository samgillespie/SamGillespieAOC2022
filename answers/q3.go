package answers

func Day3() []interface{} {
	data := ReadInputAsStr(3)
	return []interface{}{q3part1(data), q3part2(data)}
}

func findOverlap(a []byte, b []byte) byte {
	for _, i := range a {
		for _, j := range b {
			if i == j {
				return i
			}
		}
	}
	return '!'
}

func findAllOverlap(a []byte, b []byte) []byte {
	solution := []byte{}
	for _, i := range a {
		for _, j := range b {
			if i == j {
				solution = append(solution, i)
			}
		}
	}
	return solution
}

func byteToPriority(a byte) int {
	if a <= 90 {
		return int(a) - 64 + 26
	} else {
		return int(a) - 96
	}
}

func q3part1(data []string) int {
	answer := 0
	for _, row := range data {
		halfway := int(len(row) / 2)
		first_half := row[0:halfway]
		second_half := row[halfway:]
		solution := findOverlap([]byte(first_half), []byte(second_half))
		answer += byteToPriority(solution)
	}
	return answer
}

func q3part2(data []string) int {
	var row1, row2 []byte
	answer := 0
	for _, row := range data {
		if len(row1) == 0 {
			row1 = []byte(row)
		} else if len(row2) == 0 {
			row2 = []byte(row)
		} else {
			overlap := findAllOverlap(row1, row2)
			solution := findAllOverlap(overlap, []byte(row))
			answer += byteToPriority(solution[0])
			row1 = []byte{}
			row2 = []byte{}
		}
	}
	return answer
}
