package answers

import "fmt"

func Day12() []interface{} {
	data := ReadInputAsStr(12)
	grid, start := ParseGrid(data)
	return []interface{}{q12part1(grid, start), q12part2(grid)}
}

func ParseGrid(data []string) ([]*GridCell, *GridCell) {
	result := []*GridCell{}
	cellIndex := map[Vector]*GridCell{}
	var startCell *GridCell

	// Create the cells
	for y, row := range data {
		for x, char := range row {
			height := char - 97 // a=0, b=1, etc.
			pos := Vector{x: x, y: y}
			cell := GridCell{
				pos:      pos,
				height:   int(height),
				adjacent: []*GridCell{},
			}
			if height == -14 {
				startCell = &cell
				cell.height = 0
			}
			if height == -28 {
				cell.height = 25
				cell.target = true
			}
			result = append(result, &cell)
			cellIndex[pos] = &cell
		}
	}

	// Now build a map of all the cells
	for y, row := range data {
		for x, _ := range row {
			cell := cellIndex[Vector{x: x, y: y}]
			adj := cell.adjacent
			if x > 0 {
				left := cellIndex[Vector{x: x - 1, y: y}]
				if left.height-cell.height <= 1 {
					adj = append(adj, left)
				}
			}
			if x < len(row)-1 {
				right := cellIndex[Vector{x: x + 1, y: y}]
				if right.height-cell.height <= 1 {
					adj = append(adj, right)
				}
			}
			if y > 0 {
				up := cellIndex[Vector{x: x, y: y - 1}]
				if up.height-cell.height <= 1 {
					adj = append(adj, up)
				}
			}
			if y < len(data)-1 {
				down := cellIndex[Vector{x: x, y: y + 1}]
				if down.height-cell.height <= 1 {
					adj = append(adj, down)
				}
			}
			cell.adjacent = adj
		}
	}
	return result, startCell
}

type GridCell struct {
	pos      Vector
	height   int
	adjacent []*GridCell
	target   bool
}

func PrintCells(cells []*GridCell) {
	for _, cell := range cells {
		fmt.Println(cell.pos, cell.height)
	}
}

func FindDistance(grid []*GridCell, start *GridCell) int {
	visited := map[Vector]int{start.pos: 0}
	walkerHeads := []*GridCell{start}
	steps := 0
	for len(walkerHeads) > 0 {
		nextStep := []*GridCell{}
		for _, head := range walkerHeads {
			for _, destination := range head.adjacent {
				// Check to see if we've already made it to this position in another way
				_, reached := visited[destination.pos]
				if reached {
					continue
				}
				if destination.target == true {
					return steps + 1
				}
				nextStep = append(nextStep, destination)
				visited[destination.pos] = steps + 1
			}
		}
		walkerHeads = nextStep
		steps++
	}
	return 999999
}

func q12part1(grid []*GridCell, start *GridCell) int {
	// Walk the space
	return FindDistance(grid, start)
}

func q12part2(grid []*GridCell) int {
	startPoints := []*GridCell{}
	for _, cell := range grid {
		if cell.height == 0 {
			startPoints = append(startPoints, cell)
		}
	}
	min_val := 999999999
	for _, start := range startPoints {
		dist := FindDistance(grid, start)
		if dist < min_val {
			min_val = dist
		}
	}

	return min_val
}
