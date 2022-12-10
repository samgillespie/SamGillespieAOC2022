package answers

import (
	"fmt"
	"strconv"
	"strings"
)

func Day10() []interface{} {
	data := ReadInputAsStr(10)
	return []interface{}{q10part1(data), q10part2(data)}
}

func q10part1(data []string) int {
	X := 1
	cycle := 0
	toAdd := 0
	cursor := 0
	signalStrength := 0
	for {
		if cycle == 20-1 || cycle == 60-1 || cycle == 100-1 || cycle == 140-1 || cycle == 180-1 || cycle == 220-1 {
			signalStrength += (cycle + 1) * X
		}
		cycle++
		if toAdd != 0 {
			X += toAdd
			toAdd = 0
			continue
		}
		if cursor >= len(data) {
			break
		}
		instruction := data[cursor]
		cursor++
		if instruction == "noop" {
			continue
		}
		split := strings.Split(instruction, " ")
		toAdd, _ = strconv.Atoi(split[1])
	}

	return signalStrength
}

func PrintMessage(data [][]rune) {
	for _, row := range data {
		fmt.Println(string(row))
	}
}

func q10part2(data []string) []string {
	image := [][]rune{}
	for i := 0; i < 6; i++ {
		image = append(image, make([]rune, 40))
	}
	X := 1
	cycle := 0
	toAdd := 0
	cursor := 0
	for {
		colPosition := cycle / 40
		rowPosition := cycle % 40
		if colPosition == 6 {
			break
		}
		cycle++
		if rowPosition >= X-1 && rowPosition <= X+1 {
			image[colPosition][rowPosition] = '#'
		} else {
			image[colPosition][rowPosition] = '.'
		}
		if toAdd != 0 {
			X += toAdd
			toAdd = 0
			continue
		}
		if cursor >= len(data) {
			break
		}
		instruction := data[cursor]
		cursor++
		if instruction == "noop" {
			continue
		}
		split := strings.Split(instruction, " ")
		toAdd, _ = strconv.Atoi(split[1])
	}
	// PrintMessage(image)
	result := []string{}
	for _, row := range image {
		result = append(result, string(row)+"\n")
	}
	result[0] = "\n " + result[0]
	return result
}
