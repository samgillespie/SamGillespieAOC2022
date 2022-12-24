package answers

import "fmt"

func Day24() []interface{} {
	data := ReadInputAsStr(24)
	return []interface{}{q24part1(data), q24part2(data)}
}

type Blizzard struct {
	position  Vector
	direction Vector
}

type BlizzardList []Blizzard

func (bl BlizzardList) GetNextPositions(walkable map[Vector]bool, bounds VectorBounds) ([]Vector, BlizzardList) {
	nextPositions := []Vector{}
	nextPositionsMap := map[Vector]bool{}
	for pos, blizzard := range bl {
		nextPosition := blizzard.position.Add(blizzard.direction)
		isValid, notFound := walkable[nextPosition]
		if !notFound {
			panic("Shouldn't happen")
		}
		if isValid == false {
			if blizzard.direction.y == 1 {
				nextPosition.y = 1
			} else if blizzard.direction.y == -1 {
				nextPosition.y = bounds.ymax - 2
			} else if blizzard.direction.x == -1 {
				nextPosition.x = bounds.xmax - 2
			} else if blizzard.direction.x == 1 {
				nextPosition.x = 1
			}
		}
		blizzard.position = nextPosition
		bl[pos] = blizzard
		nextPositionsMap[blizzard.position] = true
	}
	for next := range nextPositionsMap {
		nextPositions = append(nextPositions, next)
	}
	return nextPositions, bl
}

func CopyMapVectorBool(input map[Vector]bool) map[Vector]bool {
	new := map[Vector]bool{}
	for key, value := range input {
		new[key] = value
	}
	return new
}

func ParseBlizzardInput(data []string) (map[Vector]bool, []Blizzard) {
	blizzards := []Blizzard{}
	walkable := map[Vector]bool{}
	for y, row := range data {
		for x, cell := range row {
			if cell == '#' {
				walkable[Vector{x: x, y: y}] = false
			} else {
				walkable[Vector{x: x, y: y}] = true
			}

			if cell == '>' {
				blizzard := Blizzard{position: Vector{x: x, y: y}, direction: Vector{x: 1, y: 0}}
				blizzards = append(blizzards, blizzard)
			} else if cell == '<' {
				blizzard := Blizzard{position: Vector{x: x, y: y}, direction: Vector{x: -1, y: 0}}
				blizzards = append(blizzards, blizzard)
			} else if cell == 'v' {
				blizzard := Blizzard{position: Vector{x: x, y: y}, direction: Vector{x: 0, y: 1}}
				blizzards = append(blizzards, blizzard)
			} else if cell == '^' {
				blizzard := Blizzard{position: Vector{x: x, y: y}, direction: Vector{x: 0, y: -1}}
				blizzards = append(blizzards, blizzard)
			}
		}
	}
	return walkable, blizzards
}

type CurrentBlizzardStep struct {
	time     int
	position Vector
	bounds   *VectorBounds
	target   *Vector
}

func (cbs CurrentBlizzardStep) Copy() CurrentBlizzardStep {
	return CurrentBlizzardStep{
		time:     cbs.time,
		position: cbs.position,
		bounds:   cbs.bounds,
		target:   cbs.target,
	}
}

func (cbs CurrentBlizzardStep) GetNextSteps(walkable map[Vector]bool) []CurrentBlizzardStep {
	//South
	exits := []Vector{
		{x: cbs.position.x, y: cbs.position.y + 1}, // South
		{x: cbs.position.x, y: cbs.position.y - 1}, // North
		{x: cbs.position.x - 1, y: cbs.position.y}, // West
		{x: cbs.position.x + 1, y: cbs.position.y}, // East
		cbs.position, // Stay
	}
	nextSteps := []CurrentBlizzardStep{}
	for _, target := range exits {
		isWalkable, exists := walkable[target]
		if exists == false || isWalkable == false {
			continue
		}
		new := cbs.Copy()
		new.position = target
		nextSteps = append(nextSteps, new)
	}
	return nextSteps
}

func PruneBlizzardSteps(steps []CurrentBlizzardStep) []CurrentBlizzardStep {
	toKeep := []CurrentBlizzardStep{}
	valid := map[Vector]CurrentBlizzardStep{}
	// Since we are running BFS, where all are at the same step, we can just take any instance
	for _, step := range steps {
		valid[step.position] = step
	}
	for _, step := range valid {
		toKeep = append(toKeep, step)
	}
	return toKeep
}

func NextStep(walkable map[Vector]bool, blizzards BlizzardList, current []CurrentBlizzardStep) ([]CurrentBlizzardStep, BlizzardList) {
	blizzardPositions, blizzards := blizzards.GetNextPositions(walkable, *current[0].bounds)
	for _, blizzardPosition := range blizzardPositions {
		// Since walkable is passed via copy, we can overwrite it and be fine here
		walkable[blizzardPosition] = false
	}
	nextSteps := []CurrentBlizzardStep{}
	for _, step := range current {
		iteration := step.GetNextSteps(walkable)
		nextSteps = append(nextSteps, iteration...)
	}
	// PrintMap(walkable, *current[0].bounds)
	return nextSteps, blizzards
}

func PrintMap(walkable map[Vector]bool, bounds VectorBounds) {
	for y := 0; y < bounds.ymax; y++ {
		row := ""
		for x := 0; x < bounds.xmax; x++ {
			isValid, exists := walkable[Vector{x: x, y: y}]
			if isValid == false {
				row += "#"
			} else if exists == false {
				row += "!"
			} else {
				row += "."
			}
		}
		fmt.Println(row)
	}
}

func q24part1(data []string) int {
	walkable, blizzards := ParseBlizzardInput(data)
	bounds := VectorBounds{xmin: 0, ymin: 0, xmax: len(data[0]), ymax: len(data)}
	target := Vector{x: bounds.xmax - 2, y: bounds.ymax - 1}

	currentSteps := []CurrentBlizzardStep{{time: 0, position: Vector{x: 1, y: 0}, target: &target, bounds: &bounds}}
	for {
		walkCopy := CopyMapVectorBool(walkable)
		currentSteps, blizzards = NextStep(walkCopy, blizzards, currentSteps)
		for pos, step := range currentSteps {
			step.time++
			if step.position.Equals(step.target) {
				return step.time
			}
			currentSteps[pos] = step
		}
		currentSteps = PruneBlizzardSteps(currentSteps)
		if currentSteps[0].time%1000 == 0 {
			fmt.Printf("At time %d have %d positions %d\n", currentSteps[0].time, len(currentSteps), currentSteps[0].time%1000)
		}
	}
}

func q24part2(data []string) int {
	walkable, blizzards := ParseBlizzardInput(data)
	bounds := VectorBounds{xmin: 0, ymin: 0, xmax: len(data[0]), ymax: len(data)}
	targets := []Vector{{x: 1, y: 0}, {x: bounds.xmax - 2, y: bounds.ymax - 1}}
	startTarget := Vector{x: bounds.xmax - 2, y: bounds.ymax - 1}
	currentSteps := []CurrentBlizzardStep{{time: 0, position: Vector{x: 1, y: 0}, target: &startTarget, bounds: &bounds}}
	for {
		walkCopy := CopyMapVectorBool(walkable)
		currentSteps, blizzards = NextStep(walkCopy, blizzards, currentSteps)
		bailout := CurrentBlizzardStep{}
		for pos, step := range currentSteps {
			step.time++
			if step.position.Equals(step.target) {
				bailout = step.Copy()
				if len(targets) == 2 {
					bailout.target = &targets[0]
					targets = targets[1:]
				} else if len(targets) == 1 {
					bailout.target = &targets[0]
					targets = []Vector{}
				} else {
					return step.time
				}
				break
			}

			currentSteps[pos] = step
		}
		if bailout.time != 0 {
			currentSteps = []CurrentBlizzardStep{bailout}
		}
		currentSteps = PruneBlizzardSteps(currentSteps)
		if currentSteps[0].time%1000 == 0 {
			fmt.Printf("At time %d have %d positions %d\n", currentSteps[0].time, len(currentSteps), currentSteps[0].time%1000)
		}
	}
}
