package answers

import "fmt"

type MonkeyMapCell struct {
	x        int
	y        int
	walkable bool
	up       *MonkeyMapCell
	down     *MonkeyMapCell
	left     *MonkeyMapCell
	right    *MonkeyMapCell

	// If we pass over a threshold, we may need to change our orientation
	upRotate    Vector
	downRotate  Vector
	leftRotate  Vector
	rightRotate Vector
}

func (mmc MonkeyMapCell) Validate() {
	if mmc.up == nil {
		fmt.Printf("Validation failed for %d, %d up\n", mmc.x, mmc.y)
	}
	if mmc.down == nil {
		fmt.Printf("Validation failed for %d, %d down\n", mmc.x, mmc.y)
	}
	if mmc.left == nil {
		fmt.Printf("Validation failed for %d, %d left\n", mmc.x, mmc.y)
	}
	if mmc.right == nil {
		fmt.Printf("Validation failed for %d, %d right\n", mmc.x, mmc.y)
	}
}

func (mmc *MonkeyMapCell) Move(direction Vector) (*MonkeyMapCell, Vector, bool) {

	if direction.y == -1 {
		target := mmc.up
		if mmc.up.walkable == false {
			return mmc, direction, false
		}
		if mmc.upRotate.x != 0 || mmc.upRotate.y != 0 {
			direction = mmc.upRotate
		}
		return target, direction, true
	}

	if direction.y == 1 {
		target := mmc.down
		if mmc.down.walkable == false {
			return mmc, direction, false
		}
		if mmc.downRotate.x != 0 || mmc.downRotate.y != 0 {
			direction = mmc.downRotate
		}
		return target, direction, true
	}

	if direction.x == 1 {
		target := mmc.right
		if mmc.right.walkable == false {
			return mmc, direction, false
		}
		if mmc.rightRotate.x != 0 || mmc.rightRotate.y != 0 {
			direction = mmc.rightRotate
		}
		return target, direction, true
	}

	if direction.x == -1 {
		target := mmc.left
		if mmc.left.walkable == false {
			return mmc, direction, false
		}
		if mmc.leftRotate.x != 0 || mmc.leftRotate.y != 0 {
			direction = mmc.leftRotate
		}
		return target, direction, true
	}
	panic("How'd you get here?")
}

func ParseMonkeyMapIntoCells(data []string) map[Vector]*MonkeyMapCell {
	monkeyMap := map[Vector]*MonkeyMapCell{}
	for y, row := range data[0 : len(data)-2] {
		for x, val := range row {
			if val == ' ' {
				continue
			}
			cell := MonkeyMapCell{
				x:        x,
				y:        y,
				walkable: val == '.',
			}
			monkeyMap[Vector{x: x, y: y}] = &cell
		}
	}

	// Do the regular connections
	for pos, cell := range monkeyMap {
		leftPos := Vector{x: pos.x - 1, y: pos.y}
		leftCell, leftExists := monkeyMap[leftPos]
		if leftExists {
			cell.left = leftCell
		}
		rightPos := Vector{x: pos.x + 1, y: pos.y}
		rightCell, rightExists := monkeyMap[rightPos]
		if rightExists {
			cell.right = rightCell
		}
		upPos := Vector{x: pos.x, y: pos.y - 1}
		upCell, upExists := monkeyMap[upPos]
		if upExists {
			cell.up = upCell
		}
		downPos := Vector{x: pos.x, y: pos.y + 1}
		downCell, downExists := monkeyMap[downPos]
		if downExists {
			cell.down = downCell
		}
	}

	// Do the vertices.
	// Little cheatsheet to map the verticies.
	// Do #1
	for i := 0; i < 50; i++ {
		pos_a, exists := monkeyMap[Vector{x: i + 100, y: 49}]
		if !exists {
			panic("error in #1a")
		}
		pos_b, exists := monkeyMap[Vector{x: 99, y: 50 + i}]
		if !exists {
			panic("error in #1b")
		}
		pos_a.down = pos_b
		pos_a.downRotate = Vector{x: -1, y: 0}
		pos_b.right = pos_a
		pos_b.rightRotate = Vector{x: 0, y: -1}
	}

	// Do #2
	for i := 0; i < 50; i++ {
		pos_a, exists := monkeyMap[Vector{x: 149, y: i}]
		if !exists {
			panic("error in #1a")
		}
		pos_b, exists := monkeyMap[Vector{x: 99, y: 149 - i}]
		if !exists {
			panic("error in #1b")
		}
		pos_a.right = pos_b
		pos_a.rightRotate = Vector{x: -1, y: 0}
		pos_b.right = pos_a
		pos_b.rightRotate = Vector{x: -1, y: 0}
	}

	// Do #3
	for i := 0; i < 50; i++ {
		pos_a, exists := monkeyMap[Vector{x: i + 100, y: 0}]
		if !exists {
			panic("error in #3a")
		}
		pos_b, exists := monkeyMap[Vector{x: i, y: 199}]
		if !exists {
			panic("error in #3b")
		}
		pos_a.up = pos_b
		pos_b.down = pos_a
		// No rotation issues here
	}

	// Do #9
	for i := 0; i < 50; i++ {
		pos_a, exists := monkeyMap[Vector{x: 50 + i, y: 149}]
		if !exists {
			panic("error in #9a")
		}
		pos_b, exists := monkeyMap[Vector{x: 49, y: 150 + i}]
		if !exists {
			panic("error in #9b")
		}
		pos_a.down = pos_b
		pos_a.downRotate = Vector{x: -1, y: 0}
		pos_b.right = pos_a
		pos_b.rightRotate = Vector{x: 0, y: -1}
	}

	// Do #alpha
	for i := 0; i < 50; i++ {
		pos_a, exists := monkeyMap[Vector{x: 50, y: 50 + i}]
		if !exists {
			panic("error in #alphaa")
		}
		pos_b, exists := monkeyMap[Vector{x: i, y: 100}]
		if !exists {
			panic("error in #alphab")
		}
		pos_a.left = pos_b
		pos_a.leftRotate = Vector{x: 0, y: 1}
		pos_b.up = pos_a
		pos_b.upRotate = Vector{x: 1, y: 0}
	}

	// Do #beta
	for i := 0; i < 50; i++ {
		pos_a, exists := monkeyMap[Vector{x: 50, y: i}]
		if !exists {
			panic("error in #betaa")
		}
		pos_b, exists := monkeyMap[Vector{x: 0, y: 149 - i}]
		if !exists {
			panic("error in #betab")
		}
		pos_a.left = pos_b
		pos_a.leftRotate = Vector{x: 1, y: 0}
		pos_b.left = pos_a
		pos_b.leftRotate = Vector{x: 1, y: 0}
	}

	// Do #pi
	for i := 0; i < 50; i++ {
		pos_a, exists := monkeyMap[Vector{x: 50 + i, y: 0}]
		if !exists {
			panic("error in #pia")
		}
		pos_b, exists := monkeyMap[Vector{x: 0, y: 150 + i}]
		if !exists {
			panic("error in #pib")
		}
		pos_a.up = pos_b
		pos_a.upRotate = Vector{x: 1, y: 0}
		pos_b.left = pos_a
		pos_b.leftRotate = Vector{x: 0, y: 1}
	}

	// Validate
	for _, cell := range monkeyMap {
		cell.Validate()
	}

	return monkeyMap
}

func RunTests(monkeyMap map[Vector]*MonkeyMapCell) {

}

func q22part2(data []string) int {
	monkeyMap := ParseMonkeyMapIntoCells(data)
	instructions := ParseMonkeyInstruction(data[len(data)-1])
	currentPosition := monkeyMap[Vector{x: 50, y: 0}]
	directionList := GetDirectionList()
	currentDirection := directionList[0]

	for _, instruction := range instructions {
		if instruction.rotateLeft == true {
			currentDirection = directionList.RotateLeft(currentDirection)
			// fmt.Println("Rotated left")
			continue
		}
		if instruction.rotateRight == true {
			currentDirection = directionList.RotateRight(currentDirection)
			// fmt.Println("Rotated right")
			continue
		}

		for i := 0; i < instruction.distance; i++ {
			newPosition, newDirection, walkable := currentPosition.Move(currentDirection)
			// fmt.Println("Moved from position", currentPosition, " to ", newPosition)
			currentPosition = newPosition
			currentDirection = newDirection
			if walkable == false {
				// fmt.Println("Hit A Wall")
				break
			}
		}
	}

	solution := (currentPosition.y+1)*1000 + (currentPosition.x+1)*4
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
