package answers

import (
	"strconv"
	"strings"
)

func Day8() []interface{} {
	data := ReadInputAsStr(8)
	return []interface{}{q8part1(data), q8part2(data)}

}

type TreeGrid [][]int
type TreeVisible [][]bool

func (tv TreeVisible) VisibleTrees() int {
	counter := 0
	for _, row := range tv {
		for _, cell := range row {
			if cell {
				counter++
			}
		}
	}
	return counter
}

func (tg TreeGrid) HorizSlice(rowidx int, visibilityGrid TreeVisible, reverse bool) TreeVisible {
	vis := -1
	for i := 0; i < len(tg[rowidx]); i++ {
		colidx := i
		if reverse == true {
			colidx = len(tg[rowidx]) - i - 1
		}
		cell := tg[rowidx][colidx]

		if cell > vis {
			vis = cell
			visibilityGrid[rowidx][colidx] = true
			if cell == 9 {
				break
			}
		}
	}
	return visibilityGrid
}

func (tg TreeGrid) VertSlice(colidx int, visibilityGrid TreeVisible, reverse bool) TreeVisible {
	vis := -1
	for i := 0; i < len(tg); i++ {
		rowidx := i
		if reverse == true {
			rowidx = len(tg) - i - 1
		}
		cell := tg[rowidx][colidx]

		if cell > vis {
			vis = cell
			visibilityGrid[rowidx][colidx] = true
		}
	}
	return visibilityGrid
}

func (tg TreeGrid) ScenicScore(tree_x int, tree_y int) int {
	treeHeight := tg[tree_x][tree_y]

	// Look Right
	score := 1
	vis := -1
	counter := 0
	for i := tree_x + 1; i < len(tg[0]); i++ {
		cell := tg[i][tree_y]
		vis = cell
		counter += 1
		if vis >= treeHeight {
			break
		}

	}
	score *= counter

	// Look Left
	vis = -1
	counter = 0
	for i := tree_x - 1; i >= 0; i-- {
		cell := tg[i][tree_y]

		vis = cell
		counter += 1
		if vis >= treeHeight {
			break
		}
	}
	score *= counter

	// Look Up
	vis = -1
	counter = 0
	for j := tree_y - 1; j >= 0; j-- {
		cell := tg[tree_x][j]
		vis = cell
		counter += 1
		if vis >= treeHeight {
			break
		}

	}
	score *= counter

	// Look Down
	vis = -1
	counter = 0
	for j := tree_y + 1; j < len(tg); j++ {
		cell := tg[tree_x][j]
		vis = cell
		counter += 1
		if vis >= treeHeight {
			break
		}

	}
	score *= counter
	return score
}

func ParseDay8(data []string) TreeGrid {
	result := TreeGrid{}
	for _, row := range data {
		split := strings.Split(row, "")
		intList := []int{}
		for _, num := range split {
			val, _ := strconv.Atoi(num)
			intList = append(intList, val)
		}
		result = append(result, intList)
	}
	return result
}

func MakeTreeVisible(x int, y int) TreeVisible {
	vis := TreeVisible{}
	for i := 0; i < y; i++ {
		vis = append(vis, make([]bool, x))
	}
	return vis
}

func q8part1(data []string) int {
	treeGrid := ParseDay8(data)
	treeVisible := MakeTreeVisible(len(data), len(data[0]))
	for rowidx, _ := range treeGrid {
		treeVisible = treeGrid.HorizSlice(rowidx, treeVisible, true)
		treeVisible = treeGrid.HorizSlice(rowidx, treeVisible, false)
	}
	for colidx, _ := range treeGrid[0] {
		treeVisible = treeGrid.VertSlice(colidx, treeVisible, true)
		treeVisible = treeGrid.VertSlice(colidx, treeVisible, false)
	}
	return treeVisible.VisibleTrees()
}

func q8part2(data []string) int {
	treeGrid := ParseDay8(data)

	// Skip first and last cols and rows, since their score will be 0
	winning_score := 0

	for x := 1; x < len(treeGrid)-1; x++ {
		for y := 1; y < len(treeGrid[0])-1; y++ {
			score := treeGrid.ScenicScore(x, y)
			if score > winning_score {
				winning_score = score
			}
		}
	}
	return winning_score
}
