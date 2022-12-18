package answers

import (
	"fmt"
	"strconv"
	"strings"
)

type CubeContents int

const (
	AIR          CubeContents = 1
	OBSIDIAN     CubeContents = 2
	INTERNAL_AIR CubeContents = 3
	EXTERNAL_AIR CubeContents = 4
)

func Day18() []interface{} {
	data := ReadInputAsStr(18)
	lava := ParseLava(data)
	return []interface{}{q18part1(lava), q18part2(lava)}
}

func ParseLava(data []string) []Vector3 {
	result := []Vector3{}
	for _, row := range data {
		split := strings.Split(row, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])
		result = append(result, Vector3{x: x, y: y, z: z})
	}
	return result
}

func GetNeighbours(cell Vector3) []Vector3 {
	return []Vector3{
		cell.Up(1),
		cell.Down(1),
		cell.Left(1),
		cell.Right(1),
		cell.Forward(1),
		cell.Back(1),
	}
}

func q18part1(lava []Vector3) int {
	lavaExistsMap := map[Vector3]bool{}
	for _, cell := range lava {
		lavaExistsMap[cell] = true
	}

	totalSides := 0
	for _, cell := range lava {
		neighbours := GetNeighbours(cell)
		for _, neighbour := range neighbours {
			if _, ok := lavaExistsMap[neighbour]; !ok {
				totalSides++
			}
		}
	}

	return totalSides
}

func RecursivelyMapAir(startPosition Vector3, space map[Vector3]CubeContents, bounds Vector3Bounds) map[Vector3]CubeContents {
	positionsVisited := []Vector3{}
	toProcess := []Vector3{startPosition}
	var cell Vector3
	isExternal := false
	for len(toProcess) > 0 {
		cell, toProcess = toProcess[0], toProcess[1:]
		positionsVisited = append(positionsVisited, cell)
		neighbours := GetNeighbours(cell)
		for _, neighbour := range neighbours {
			// Don't add a position, if we've already visited it
			if Vector3InSlice(neighbour, positionsVisited) {
				continue
			}

			spaceType, exists := space[neighbour]
			if exists == false {
				// We are out of bounds, continue
				continue
			}
			if spaceType == OBSIDIAN {
				continue
			} else if spaceType == EXTERNAL_AIR {
				isExternal = true
			} else {
				if AtEdge(neighbour, bounds) == true {
					isExternal = true
				}

				if Vector3InSlice(neighbour, toProcess) == false {
					toProcess = append(toProcess, neighbour)
				}
			}

		}
	}

	for _, positionsVisited := range positionsVisited {
		if isExternal {
			space[positionsVisited] = EXTERNAL_AIR
		} else {
			space[positionsVisited] = INTERNAL_AIR
		}
	}
	return space
}

func AtEdge(air Vector3, bounds Vector3Bounds) bool {
	if air.x == bounds.xmin || air.x == bounds.xmax {
		return true
	}
	if air.y == bounds.ymin || air.y == bounds.ymax {
		return true
	}
	if air.z == bounds.zmin || air.z == bounds.zmax {
		return true
	}
	return false
}

func Print(space map[Vector3]CubeContents, bounds Vector3Bounds) {
	cube := [][]string{}
	for x := bounds.xmin; x <= bounds.xmax; x++ {
		yrow := []string{}
		for y := bounds.ymin; y <= bounds.ymax; y++ {
			zstring := ""
			for z := bounds.zmin; z <= bounds.zmax; z++ {
				contents := space[Vector3{x, y, z}]
				if contents == AIR {
					zstring += "."
				} else if contents == OBSIDIAN {
					zstring += "#"
				} else if contents == INTERNAL_AIR {
					zstring += "x"
				} else if contents == EXTERNAL_AIR {
					zstring += "o"
				}
			}
			yrow = append(yrow, zstring)
		}
		cube = append(cube, yrow)
	}

	for z, slice := range cube {
		fmt.Printf("\n--------- SLICE z=%d -----------\n", z)
		for _, i := range slice {
			fmt.Println(i)
		}
	}

}

func q18part2(lava []Vector3) int {
	bounds := CalculateVector3Bounds(lava)
	// Prepare the space with values.
	space := map[Vector3]CubeContents{}
	for x := bounds.xmin; x <= bounds.xmax; x++ {
		for y := bounds.ymin; y <= bounds.ymax; y++ {
			for z := bounds.zmin; z <= bounds.zmax; z++ {
				vec := Vector3{x: x, y: y, z: z}
				space[vec] = AIR // No idea what's here yet
			}
		}
	}

	for _, cell := range lava {
		space[cell] = OBSIDIAN
	}

	for position, spaceType := range space {
		if spaceType == AIR {
			space = RecursivelyMapAir(position, space, bounds)
		}
	}

	//Print(space, bounds)

	totalSides := 0
	for _, cell := range lava {
		neighbours := GetNeighbours(cell)
		for _, neighbour := range neighbours {
			spaceType, exists := space[neighbour]
			if exists == false || spaceType == EXTERNAL_AIR {
				totalSides++
			}
		}
	}
	return totalSides
}
