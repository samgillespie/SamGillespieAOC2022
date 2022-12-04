package answers

import (
	"strconv"
	"strings"
)

func Day4() []int {
	data := ReadInputAsStr(4)
	elves := ParseElves(data)
	return []int{q4part1(elves), q4part2(elves)}
}

type Elf struct {
	start int
	end   int
}

func ParseElves(data []string) [][]Elf {
	parsed := [][]Elf{}
	for _, elf := range data {
		split := strings.Split(elf, ",")
		parsed = append(parsed, []Elf{
			ElfParse(split[0]),
			ElfParse(split[1]),
		})
	}
	return parsed
}

func ElfParse(data string) Elf {
	split := strings.Split(data, "-")
	start, _ := strconv.Atoi(split[0])
	end, _ := strconv.Atoi(split[1])
	return Elf{start: start, end: end}
}

func (e1 Elf) Contains(e2 Elf) bool {
	if e1.start >= e2.start && e1.end <= e2.end {
		return true
	}
	return false
}

func (e1 Elf) Overlaps(e2 Elf) bool {
	if e1.end >= e2.start && e1.start <= e2.end {
		return true
	}
	return false
}

func q4part1(elves [][]Elf) int {
	counter := 0
	for _, elf := range elves {
		elf1 := elf[0]
		elf2 := elf[1]
		if elf1.Contains(elf2) || elf2.Contains(elf1) {
			counter += 1
		}
	}
	return counter
}

func q4part2(elves [][]Elf) int {
	counter := 0
	for _, elf := range elves {
		elf1 := elf[0]
		elf2 := elf[1]
		if elf1.Overlaps(elf2) || elf2.Overlaps(elf1) {
			counter += 1
		}
	}
	return counter
}
