package answers

func Day2() []int {
	data := ReadInputAsStr(2)
	return []int{q2part1(data), q2part2(data)}

}

func q2part1(data []string) int {
	scoremap := map[string]int{
		"A X": 4, // Rock vs Rock - tie
		"A Y": 8, // Rock Vs Paper - Win
		"A Z": 3, // Rock Vs Scissors - Loss
		"B X": 1, // Paper vs Rock - Loss
		"B Y": 5, // Paper vs Paper - Tie
		"B Z": 9, // Paper vs Scissors - Win
		"C X": 7, // Scissors vs Rock - Win
		"C Y": 2, // Scissors vs Paper = Loss
		"C Z": 6, // Scissors vs Scissors - TIe
	}
	totalScore := 0
	for _, row := range data {
		totalScore += scoremap[row]
	}
	return totalScore
}

func q2part2(data []string) int {
	scoremap := map[string]int{
		"A X": 3, // Rock vs Scissors - loss
		"A Y": 4, // Rock Vs Rock - tie
		"A Z": 8, // Rock Vs Paper - win
		"B X": 1, // Paper vs Rock - Loss
		"B Y": 5, // Paper vs Paper - Tie
		"B Z": 9, // Paper vs Scissors - Win
		"C X": 2, // Scissors vs Paper - Loss
		"C Y": 6, // Scissors vs Scissors = Tie
		"C Z": 7, // Scissors vs Rock - Win
	}
	totalScore := 0
	for _, row := range data {
		totalScore += scoremap[row]
	}
	return totalScore
}
