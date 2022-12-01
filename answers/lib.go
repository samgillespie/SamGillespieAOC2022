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
