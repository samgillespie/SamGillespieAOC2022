package answers

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadInputAsStr(value int) []string {
	data, err := ioutil.ReadFile("./inputs/q" + strconv.Itoa(value) + ".txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}

	str_values := strings.Split(string(data), "\r\n")
	return str_values
}

func ReadInputAsInt(value int) []int {
	str_values := ReadInputAsStr(value)
	ary := make([]int, len(str_values))
	for i := range ary {
		ary[i], _ = strconv.Atoi(str_values[i])
	}
	return ary
}

func ReadCSVAsInt(value int) []int {
	str_values := ReadInputAsStr(value)
	str_values = strings.Split(str_values[0], ",")
	ary := make([]int, len(str_values))
	var err error
	for i := range str_values {
		ary[i], err = strconv.Atoi(str_values[i])
		if err != nil {
			fmt.Println(err)
		}
	}
	return ary
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func abs64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func maxSlice(slice []int) (int, int) {
	// Returns position, value
	max := -99999999999
	pos := -1
	for index, elem := range slice {
		if elem > max {
			max = elem
			pos = index
		}
	}
	return pos, max
}

func minSlice(slice []int) (int, int) {
	// Returns position, value
	min := 99999999999
	pos := -1
	for index, elem := range slice {
		if elem < min {
			min = elem
			pos = index
		}
	}
	return pos, min
}

type Vector struct {
	x int
	y int
}
type VectorBounds struct {
	xmin int
	xmax int
	ymin int
	ymax int
}

func CalculateVectorBounds(vectors []Vector) VectorBounds {
	var xmax, ymax int
	xmin := 9999999
	ymin := 9999999

	for _, vec := range vectors {
		if vec.x < xmin {
			xmin = vec.x
		}
		if vec.y < ymin {
			ymin = vec.y
		}

		if vec.x > xmax {
			xmax = vec.x
		}
		if vec.y > ymax {
			ymax = vec.y
		}

	}
	return VectorBounds{xmin, xmax, ymin, ymax}
}

func (a Vector) Print() {
	fmt.Printf("x: %d, y: %d \n", a.x, a.y)
}

func (a Vector) Add(b Vector) Vector {
	return Vector{x: a.x + b.x, y: a.y + b.y}
}

type Vector3 struct {
	x int
	y int
	z int
}

type Vector3Bounds struct {
	xmin int
	xmax int
	ymin int
	ymax int
	zmin int
	zmax int
}

func (v Vector3) Up(distance int) Vector3 {
	return Vector3{x: v.x, y: v.y + distance, z: v.z}
}

func (v Vector3) Down(distance int) Vector3 {
	return Vector3{x: v.x, y: v.y - distance, z: v.z}
}

func (v Vector3) Left(distance int) Vector3 {
	return Vector3{x: v.x + distance, y: v.y, z: v.z}
}

func (v Vector3) Right(distance int) Vector3 {
	return Vector3{x: v.x - distance, y: v.y, z: v.z}
}

func (v Vector3) Forward(distance int) Vector3 {
	return Vector3{x: v.x, y: v.y, z: v.z + distance}
}

func (v Vector3) Back(distance int) Vector3 {
	return Vector3{x: v.x, y: v.y, z: v.z - distance}
}

func (a Vector3) Print() {
	fmt.Printf("x: %d, y: %d  z: %d\n", a.x, a.y, a.z)
}

func CalculateVector3Bounds(vectors []Vector3) Vector3Bounds {
	var xmax, ymax, zmax int
	xmin := 9999999
	ymin := 9999999
	zmin := 9999999
	for _, vec := range vectors {
		if vec.x < xmin {
			xmin = vec.x
		}
		if vec.y < ymin {
			ymin = vec.y
		}
		if vec.z < zmin {
			zmin = vec.z
		}
		if vec.x > xmax {
			xmax = vec.x
		}
		if vec.y > ymax {
			ymax = vec.y
		}
		if vec.z > zmax {
			zmax = vec.z
		}
	}
	return Vector3Bounds{xmin, xmax, ymin, ymax, zmin, zmax}
}

func IntInSlice(value int, slice []int) bool {
	for _, i := range slice {
		if value == i {
			return true
		}
	}
	return false
}

func Vector3InSlice(vec Vector3, slice []Vector3) bool {
	for _, i := range slice {
		if vec.x == i.x && vec.y == i.y && vec.z == i.z {
			return true
		}
	}
	return false
}
