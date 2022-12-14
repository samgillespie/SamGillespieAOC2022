package answers

import (
	"fmt"
	"strconv"
	"strings"
)

func Day14() []interface{} {
	data := ReadInputAsStr(14)

	return []interface{}{q14part1(data), q14part2(data)}

}

type Cave int

const (
	EMPTY Cave = 0
	WALL       = 1
	SAND       = 2
)

type CaveRules []Vector
type CaveGrid [][]Cave

func (cg CaveGrid) Print() {
	for _, yval := range cg {
		row := ""
		for _, xval := range yval {
			switch xval {
			case 0:
				row += "."
			case 1:
				row += "#"
			case 2:
				row += "o"
			}
		}
		fmt.Println(row)
	}
}

func (cg CaveGrid) RaycastY(pos int) int {
	for ycoord, y := range cg {
		if y[pos] != EMPTY {
			return ycoord
		}
	}
	return -10
}
func (cg CaveGrid) CountSand() int {
	sand := 0
	for _, y := range cg {
		for _, x := range y {
			if x == SAND {
				sand++
			}
		}
	}
	return sand
}

func Bounds(rules []CaveRules) (int, int) {
	// Return minx, maxx, miny, maxy
	maxx := 0
	maxy := 0
	for _, rule := range rules {
		for _, row := range rule {
			if row.x > maxx {
				maxx = row.x
			}
			if row.y > maxy {
				maxy = row.y
			}
		}
	}
	return maxx, maxy
}

func ApplyRuleToGrid(rules CaveRules, grid CaveGrid, ybuffer int) CaveGrid {
	position := rules[0]
	position.y += ybuffer
	for _, rule := range rules[0:] {
		rule.y += ybuffer

		grid[position.y][position.x] = WALL
		for position.x != rule.x || position.y != rule.y {
			if position.x > rule.x {
				position.x--
			} else if position.x < rule.x {
				position.x++
			}
			if position.y > rule.y {
				position.y--
			} else if position.y < rule.y {
				position.y++
			}
			grid[position.y][position.x] = WALL

		}
	}
	return grid
}

func CreateGrid(rules []CaveRules, ybuffer int, addfloor bool) (CaveGrid, int) {
	// Normalize the grid over to fit the bounds
	maxx, maxy := Bounds(rules)
	grid := CaveGrid{}
	for i := 0; i < maxy+ybuffer+3; i++ {
		grid = append(grid, make([]Cave, maxx*2))
	}

	for _, rule := range rules {
		grid = ApplyRuleToGrid(rule, grid, ybuffer)
	}

	if addfloor == true {
		for i, _ := range grid[maxy+ybuffer+2] {
			grid[maxy+ybuffer+2][i] = WALL
		}
	}
	return grid, 500
}

func ParseCaveInput(data []string) []CaveRules {
	rules := make([]CaveRules, len(data))
	for index, row := range data {
		split := strings.Split(row, " -> ")
		rule := CaveRules{}
		for _, elem := range split {
			rule_split := strings.Split(elem, ",")
			x, _ := strconv.Atoi(rule_split[0])
			y, _ := strconv.Atoi(rule_split[1])
			rule = append(rule, Vector{x: x, y: y})
		}
		rules[index] = rule
	}
	return rules
}

func AddSand(grid CaveGrid, x int) (CaveGrid, bool, int, int) {
	y := 0
	for {
		// Try to stop below
		if grid[y+1][x] == EMPTY {
			y++
			if y >= len(grid)-1 {
				return grid, true, x, y
			}
			continue
		}
		if grid[y+1][x-1] == EMPTY {
			y++
			x--
			continue
		}
		if grid[y+1][x+1] == EMPTY {
			y++
			x++
			continue
		}
		grid[y][x] = WALL
		return grid, false, x, y
	}
}

func q14part1(data []string) int {
	//return 0
	rules := ParseCaveInput(data)
	grid, sandSpawn := CreateGrid(rules, 2, false)
	var finished bool
	for {
		grid, finished, _, _ = AddSand(grid, sandSpawn)
		if finished == true {
			// grid.Print()
			break
		}
	}
	// grid.Print()
	return grid.CountSand()
}

func q14part2(data []string) int {
	ybuffer := 10

	rules := ParseCaveInput(data)
	grid, sandSpawn := CreateGrid(rules, ybuffer, true)
	var x, y int
	for i := 0; i < 100000000000; i++ {
		grid, _, x, y = AddSand(grid, sandSpawn)
		if x == sandSpawn && y == ybuffer {
			break
		}
	}
	// grid.Print()
	return grid.CountSand()
}
