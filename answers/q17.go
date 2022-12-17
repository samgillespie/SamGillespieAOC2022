package answers

import "fmt"

func Day17() []interface{} {
	data := ReadInputAsStr(17)
	return []interface{}{q17part1(data[0]), q17part2(data[0])}
}

// ####
var HLINE []Vector = []Vector{
	{0, 0},
	{1, 0},
	{2, 0},
	{3, 0},
} // Relative to left bottom

// .#.
// ###
// .#.
var CROSS []Vector = []Vector{
	{0, 1},
	{1, 1},
	{2, 1},
	{1, 0},
	{1, 2},
} // Relative to left bottom

// ..#
// ..#
// ###
var BACKL []Vector = []Vector{
	{0, 0},
	{1, 0},
	{2, 0},
	{2, 1},
	{2, 2},
} // Relative to left bottom

// #
// #
// #
// #

var VLINE []Vector = []Vector{
	{0, 0},
	{0, 1},
	{0, 2},
	{0, 3},
}

// ##
// ##
var BOX []Vector = []Vector{
	{0, 0},
	{0, 1},
	{1, 0},
	{1, 1},
}

var SEQUENCE [][]Vector = [][]Vector{
	HLINE,
	CROSS,
	BACKL,
	VLINE,
	BOX,
}

func GetShape(shapeNum int) []Vector {
	return SEQUENCE[shapeNum%5]
}

type Grid map[Vector]bool

func (g Grid) yMax() int {
	var maxy int
	for vec, _ := range g {
		if vec.y > maxy {
			maxy = vec.y
		}
	}
	return maxy
}

func (g Grid) GetAsString(shape []Vector, shapePos Vector, includeShape bool) [][]byte {
	extraLength := 10
	ymax := g.yMax()
	gridStr := make([][]byte, ymax+extraLength)
	for i := 0; i < ymax+extraLength; i++ {
		gridStr[i] = make([]byte, 7)
		for j := 0; j < 7; j++ {
			gridStr[i][j] = '.'
		}
	}
	for vec, _ := range g {
		gridStr[vec.y][vec.x] = '#'
	}
	if includeShape {
		for _, cell := range shape {
			gridStr[cell.y+shapePos.y][cell.x+shapePos.x] = '@'
		}
	}
	return gridStr
}

func (g Grid) Print(shape []Vector, shapePos Vector) {
	extraLength := 10
	ymax := g.yMax()
	gridStr := g.GetAsString(shape, shapePos, true)
	for i := ymax + extraLength - 1; i >= 0; i-- {
		fmt.Println("|" + string(gridStr[i]) + "|")
	}
}

func ApplyForce(grid Grid, shape []Vector, shapePos Vector, left bool) bool {
	if left {
		// Check to see if possible
		for _, cell := range shape {
			if (cell.x + shapePos.x - 1) < 0 {
				// Can't move, we're going too far left
				return false
			}
			newCellPos := Vector{x: cell.x + shapePos.x - 1, y: cell.y + shapePos.y}
			_, isBlock := grid[newCellPos]
			if isBlock == true {
				return false
			}
		}
		return true
	} else { // right
		// Check to see if possible
		for _, cell := range shape {
			if (cell.x + shapePos.x + 1) > 6 {
				// Can't move, we're going too far right
				return false
			}
			newCellPos := Vector{x: cell.x + shapePos.x + 1, y: cell.y + shapePos.y}
			_, isBlock := grid[newCellPos]
			if isBlock == true {
				return false
			}
		}
		return true
	}
}

func ApplyDrop(grid Grid, shape []Vector, shapePos Vector) bool {
	for _, cell := range shape {
		newCellPos := Vector{x: cell.x + shapePos.x, y: cell.y + shapePos.y - 1}
		_, isBlock := grid[newCellPos]
		if isBlock == true {
			return false
		}
	}
	return true
}

func SimulateStep(grid Grid, jetPattern string, maxy int, shape []Vector, jetCursor int) (Grid, int, int) {

	shapePos := Vector{x: 2, y: maxy + 4}
	for {
		left := jetPattern[jetCursor] == '<'
		force_applied := ApplyForce(grid, shape, shapePos, left)
		if force_applied == true {
			if left {
				// fmt.Println(" - Jet Pushes one to the left")
				shapePos.x = shapePos.x - 1
			} else {
				// fmt.Println(" - Jet Pushes one to the Right")
				shapePos.x = shapePos.x + 1
			}
		} else {
			// fmt.Println(" - Jet Fails to Push")
		}
		jetCursor++
		if jetCursor >= len(jetPattern) {
			jetCursor = 0
		}
		// grid.Print(shape, shapePos)
		canDrop := ApplyDrop(grid, shape, shapePos)
		if canDrop == false {
			for _, cell := range shape {
				vec := Vector{x: cell.x + shapePos.x, y: cell.y + shapePos.y}
				grid[vec] = true
				if vec.y > maxy {
					maxy = vec.y
				}
			}
			break
		} else {
			shapePos.y--
		}
	}
	return grid, jetCursor, maxy
}

func q17part1(jetPattern string) int {
	grid := Grid{}
	for x := 0; x < 7; x++ {
		grid[Vector{x: x, y: 0}] = true
	}
	maxy := 0
	jetCursor := 0
	for i := 0; i < 2022; i++ {
		shape := GetShape(i)
		grid, jetCursor, maxy = SimulateStep(grid, jetPattern, maxy, shape, jetCursor)
	}
	return maxy
}

func q17part2(jetPattern string) int {
	// So we just need to find if the pattern repeats
	preload := 5000
	grid := Grid{}
	for x := 0; x < 7; x++ {
		grid[Vector{x: x, y: 0}] = true
	}
	maxy := 0
	jetCursor := 0
	// Run for 5000 shapes, then see if the pattern stabilizes
	yHeight := []int{}
	for i := 0; i < preload; i++ {
		shape := GetShape(i)
		grid, jetCursor, maxy = SimulateStep(grid, jetPattern, maxy, shape, jetCursor)
		yHeight = append(yHeight, maxy)
	}

	solution := 0
	difference := 0
	for interval := 1; interval < 5000; interval++ {
		value := yHeight[1000+interval] - yHeight[1000]
		correct := true
		for j := 0; j < 50; j++ {
			if yHeight[1000+j+interval]-yHeight[1000+j] != value {
				correct = false
			}
		}
		if correct == true {
			fmt.Print()
			solution = interval
			difference = value
			break
		}
	}

	fmt.Println(solution, difference)
	fmt.Printf("Every %d shapes has an increase in height of %d \n", solution, difference)

	totalAmount := 1000000000000 - preload
	remainder := totalAmount % solution
	backwards := solution - remainder
	start := preload - backwards
	// Go back to the neatest position in the preloaded values
	initial := yHeight[start]

	totalAmount = 1000000000000 - start
	if totalAmount%solution != 0 {
		panic("wat")
	}
	return (totalAmount/solution)*difference + initial - 1
}
