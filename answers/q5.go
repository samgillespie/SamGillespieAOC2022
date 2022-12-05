package answers

import (
	"strconv"
	"strings"
)

func Day5() []interface{} {
	data := ReadInputAsStr(5)

	cratesP1, instructions := ParseDay5(data)
	cratesP2 := make([]Crate, len(cratesP1))
	copy(cratesP2, cratesP1) // Copy for part 1 and part 2
	return []interface{}{q5part1(cratesP1, instructions), q5part2(cratesP2, instructions)}
}

var CRATEPILES int = 9

type Instruction struct {
	number int
	origin int
	target int
}

type Crate []string

func DoInstructionP1(inst Instruction, crates []Crate) []Crate {
	for i := 0; i < inst.number; i++ {
		crate, newOrigin := crates[inst.origin-1][0], crates[inst.origin-1][1:]
		crates[inst.origin-1] = newOrigin
		crates[inst.target-1] = append([]string{crate}, crates[inst.target-1]...)
	}
	return crates
}

func DoInstructionP2(inst Instruction, crates []Crate) []Crate {
	poppedCrates, newOrigin := crates[inst.origin-1][0:inst.number], crates[inst.origin-1][inst.number:]
	copyOrigin := append([]string{}, newOrigin...) // Deference the Origin.
	newTarget := append(poppedCrates, crates[inst.target-1]...)
	crates[inst.target-1] = newTarget
	crates[inst.origin-1] = copyOrigin
	return crates
}

func ParseDay5(data []string) ([]Crate, []Instruction) {
	crates := []Crate{}
	for i := 0; i < CRATEPILES; i++ {
		crates = append(crates, Crate{})
	}
	instructions := []Instruction{}
	isCrates := true
	for _, row := range data {
		if row == "" {
			continue
		}
		if !isCrates {
			splits := strings.Split(row, " ")
			number, _ := strconv.Atoi(splits[1])
			origin, _ := strconv.Atoi(splits[3])
			target, _ := strconv.Atoi(splits[5])
			instructions = append(instructions, Instruction{
				number: number, origin: origin, target: target,
			})
		}
		if isCrates {
			for i := 0; i < CRATEPILES; i++ {
				index := i*4 + 1
				char := row[index]
				if char == '1' {
					isCrates = false
					break
				}
				if char != ' ' {
					crates[i] = append(crates[i], string(char))
				}
			}
		}
	}
	return crates, instructions
}

func q5part1(crates []Crate, instructions []Instruction) string {
	for _, instruction := range instructions {
		crates = DoInstructionP1(instruction, crates)
	}
	solution := ""
	for _, crate := range crates {
		solution += crate[0]
	}
	return solution
}

func q5part2(crates []Crate, instructions []Instruction) string {
	for _, instruction := range instructions {
		crates = DoInstructionP2(instruction, crates)
	}
	solution := ""
	for _, crate := range crates {
		solution += crate[0]
	}
	return solution
}
