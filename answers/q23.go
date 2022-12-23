package answers

func Day23() []interface{} {
	data := ReadInputAsStr(23)
	return []interface{}{q23part1(data), q23part2(data)}

}

type GroveElf struct {
	x int
	y int
}

func ParseGroveElves(data []string) map[Vector]*GroveElf {
	groveElves := map[Vector]*GroveElf{}
	for y, row := range data {
		for x, cell := range row {
			if cell == '#' {
				groveElves[Vector{x, y}] = &GroveElf{x: x, y: y}
			}
		}
	}
	return groveElves
}

func PrintGroveElves(elves map[Vector]*GroveElf, bounds VectorBounds) {
	for y := bounds.ymin; y < bounds.ymax; y++ {
		row := ""
		for x := bounds.xmin; x < bounds.xmax; x++ {
			_, exists := elves[Vector{x, y}]
			if exists {
				row += "#"
			} else {
				row += "."
			}
		}
		print(row + "\n")
	}
}

var DIRECTIONS DirectionList

// Reuse Direction List from Q22
func GetDirectionListGroveElves() DirectionList {
	return DirectionList{{x: 0, y: -1}, {x: 0, y: 1}, {x: -1, y: 0}, {x: 1, y: 0}}
}

func ProposeMove(elf Vector, groveElves map[Vector]*GroveElf, dirStart int) Vector {
	// If everywhere is elves, just stay where you are
	if dirStart >= 8 {
		return elf
	}

	_, N := groveElves[Vector{x: elf.x, y: elf.y - 1}]
	_, NE := groveElves[Vector{x: elf.x + 1, y: elf.y - 1}]
	_, NW := groveElves[Vector{x: elf.x - 1, y: elf.y - 1}]
	_, S := groveElves[Vector{x: elf.x, y: elf.y + 1}]
	_, SE := groveElves[Vector{x: elf.x + 1, y: elf.y + 1}]
	_, SW := groveElves[Vector{x: elf.x - 1, y: elf.y + 1}]
	_, W := groveElves[Vector{x: elf.x - 1, y: elf.y}]
	_, E := groveElves[Vector{x: elf.x + 1, y: elf.y}]

	//If no adjacent elves, stay still
	if !N && !NE && !NW && !S && !SE && !SW && !W && !E {
		return elf
	}

	for i := 0; i <= 3; i++ {
		dir := DIRECTIONS.GetElement(dirStart + i)
		// North
		if dir.y == -1 {

			if N || NE || NW {
				continue
			}
			return Vector{x: elf.x, y: elf.y - 1}
		}

		// South
		if dir.y == 1 {

			if S || SE || SW {
				continue
			}
			return Vector{x: elf.x, y: elf.y + 1}
		}

		// West
		if dir.x == -1 {

			if W || NW || SW {
				continue
			}
			return Vector{x: elf.x - 1, y: elf.y}
		}

		// East
		if dir.x == 1 {
			if E || NE || SE {
				continue
			}
			return Vector{x: elf.x + 1, y: elf.y}
		}
	}
	return elf
}

func q23part1(data []string) int {
	// Set global variable
	DIRECTIONS = GetDirectionListGroveElves()
	groveElves := ParseGroveElves(data)

	directionsCursor := 0

	for i := 0; i < 10; i++ {
		proposals := map[Vector]*GroveElf{}
		nullProposals := map[Vector]bool{}
		// Do Proposals
		for elfPos, elf := range groveElves {
			movement := ProposeMove(elfPos, groveElves, directionsCursor)
			_, proposalDuplicate := proposals[movement]
			if proposalDuplicate == true {
				nullProposals[movement] = true
			} else {
				proposals[movement] = elf
			}
		}

		//Cancel null proposals
		for pos := range nullProposals {
			delete(proposals, pos)
		}

		// Do Movement
		for newPos, elf := range proposals {
			origPos := Vector{x: elf.x, y: elf.y}
			delete(groveElves, origPos)
			elf.x = newPos.x
			elf.y = newPos.y
			groveElves[newPos] = elf
		}
		directionsCursor = (directionsCursor + 1) % 4
		//fmt.Println("\n\n")
		//PrintGroveElves(groveElves, groveBounds)

	}

	vectors := []Vector{}
	for elf := range groveElves {
		vectors = append(vectors, elf)
	}
	bounds := CalculateVectorBounds(vectors)

	// Bounds are out by 1 for some reason
	bounds.ymax++
	bounds.xmax++
	// PrintGroveElves(groveElves, bounds)
	cells := (bounds.xmax - bounds.xmin) * (bounds.ymax - bounds.ymin)
	return cells - len(groveElves)
}

func q23part2(data []string) int {
	DIRECTIONS = GetDirectionListGroveElves()
	groveElves := ParseGroveElves(data)
	directionsCursor := 0
	counter := 1
	for {
		proposals := map[Vector]*GroveElf{}
		nullProposals := map[Vector]bool{}
		// Do Proposals
		for elfPos, elf := range groveElves {
			movement := ProposeMove(elfPos, groveElves, directionsCursor)
			// If we are staying still continue
			if movement.x == elfPos.x && movement.y == elfPos.y {
				continue
			}

			_, proposalDuplicate := proposals[movement]
			if proposalDuplicate == true {
				nullProposals[movement] = true
			} else {
				proposals[movement] = elf
			}
		}

		//Cancel null proposals
		for pos := range nullProposals {
			delete(proposals, pos)
		}
		if len(proposals) == 0 {
			break
		}
		// Do Movement
		for newPos, elf := range proposals {
			origPos := Vector{x: elf.x, y: elf.y}
			delete(groveElves, origPos)
			elf.x = newPos.x
			elf.y = newPos.y
			groveElves[newPos] = elf
		}
		counter++
		directionsCursor = (directionsCursor + 1) % 4
	}
	return counter
}

// Too High - 6561
// Too High - 5329
// Too Low - 2605
// Wrong - 2752
