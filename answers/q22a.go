package answers

import (
	"strconv"
)

func Day22() []interface{} {
	data := ReadInputAsStr(22)
	return []interface{}{q22part1(data), q22part2(data)}
}

func ParseMonkeyInstruction(instr string) []MonkeyInstruction {

	active := ""
	instructions := []MonkeyInstruction{}
	for _, char := range instr {
		if char == 'R' || char == 'L' {
			dist, _ := strconv.Atoi(active)
			instructions = append(instructions, MonkeyInstruction{distance: dist})
			if char == 'R' {
				instructions = append(instructions, MonkeyInstruction{rotateRight: true})
			} else {
				instructions = append(instructions, MonkeyInstruction{rotateLeft: true})
			}
			active = ""
		} else {
			active += string(char)
		}
	}
	dist, _ := strconv.Atoi(active)
	instructions = append(instructions, MonkeyInstruction{distance: dist})
	return instructions
}

func ParseMonkeyMap(data []string) (MonkeyMap, []MonkeyInstruction, Vector) {
	monkeyMap := MonkeyMap{}
	monkeyMap.pos = map[Vector]bool{}
	for y, row := range data[0 : len(data)-2] {
		for x, val := range row {
			if val == ' ' {
				continue
			}
			monkeyMap.pos[Vector{x: x, y: y}] = val == '.'
		}
	}
	startingPos := Vector{x: 999, y: 0}
	for pos, val := range monkeyMap.pos {
		if pos.x > monkeyMap.maxX {
			monkeyMap.maxX = pos.x
		}
		if pos.y > monkeyMap.maxY {
			monkeyMap.maxY = pos.y
		}
		if val == true && pos.y == 0 && pos.x < startingPos.x {
			startingPos.x = pos.x
		}
	}
	instructions := ParseMonkeyInstruction(data[len(data)-1])

	return monkeyMap, instructions, startingPos
}

type DirectionList []Vector

func (dl DirectionList) RotateRight(direction Vector) Vector {
	for pos, elem := range dl {
		if elem.x == direction.x && elem.y == direction.y {
			if pos == 3 {
				return dl[0]
			}
			return dl[pos+1]
		}
	}
	panic("Rotated out of existence")
}

func (dl DirectionList) RotateLeft(direction Vector) Vector {
	for pos, elem := range dl {
		if elem.x == direction.x && elem.y == direction.y {
			if pos == 0 {
				return dl[3]
			}
			return dl[pos-1]
		}
	}
	panic("Rotated out of existence")
}

func (dl DirectionList) GetElement(pos int) Vector {
	return dl[pos%4]
}

func GetDirectionList() DirectionList {
	return DirectionList{{x: 1, y: 0}, {x: 0, y: 1}, {x: -1, y: 0}, {x: 0, y: -1}}
}

type MonkeyMap struct {
	pos  map[Vector]bool
	maxX int
	maxY int
}

func (m MonkeyMap) Move(initialPosition Vector, direction Vector) (Vector, bool) {
	// Returns a Vector of your position, and a bool if we've hit a wall
	position := initialPosition
	for {

		position = position.Add(direction)
		// Map around
		if position.x > m.maxX {
			position.x = 0
		}
		if position.y > m.maxY {
			position.y = 0
		}
		if position.x < 0 {
			position.x = m.maxX
		}
		if position.y < 0 {
			position.y = m.maxY
		}
		nextPosIsWalkable, nextPosIsSpace := m.pos[position]
		// fmt.Println(position, nextPosIsWalkable, nextPosIsSpace, direction)
		if nextPosIsSpace == false {
			continue
		}
		if nextPosIsWalkable {
			return position, false
		} else {
			return initialPosition, true
		}

	}
}

type MonkeyInstruction struct {
	distance    int
	rotateRight bool
	rotateLeft  bool
}

func q22part1(data []string) int {
	monkeyMap, instructions, position := ParseMonkeyMap(data)
	directionList := GetDirectionList()
	currentDirection := directionList[0]
	for _, instruction := range instructions {
		if instruction.rotateLeft == true {
			currentDirection = directionList.RotateLeft(currentDirection)
			continue
		}
		if instruction.rotateRight == true {
			currentDirection = directionList.RotateRight(currentDirection)
			continue
		}

		for i := 0; i < instruction.distance; i++ {
			newPosition, stopped := monkeyMap.Move(position, currentDirection)
			if stopped == true {
				// fmt.Println("Hit a Wall")
				break
			}
			// fmt.Println("Moved from position", position, " to ", newPosition)
			position = newPosition
		}
	}
	solution := (position.y+1)*1000 + (position.x+1)*4
	if currentDirection.x == 1 {
		solution += 0
	} else if currentDirection.y == 1 {
		solution += 1
	} else if currentDirection.x == -1 {
		solution += 2
	} else if currentDirection.y == -1 {
		solution += 3
	}
	return solution
}
