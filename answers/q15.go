package answers

import (
	"fmt"
	"regexp"
	"strconv"
)

type Sensor struct {
	x  int
	y  int
	bx int // Beaconx
	by int // Beacony
}
type xRange struct {
	xmin int
	xmax int
}

func (r1 xRange) Combine(r2 xRange) (xRange, bool) {
	// r1 5, 10,  r2 9, 12 -> 5, 12
	//
	// 9, 12,   5, 10 -> 5, 12
	//
	if r1.xmax >= r2.xmin && r1.xmin <= r2.xmin {
		if r1.xmax > r2.xmax {
			return xRange{xmin: r1.xmin, xmax: r1.xmax}, true
		} else {
			return xRange{xmin: r1.xmin, xmax: r2.xmax}, true
		}
	}
	if r2.xmax >= r1.xmin && r2.xmin <= r1.xmin {
		if r1.xmax > r2.xmax {
			return xRange{xmin: r2.xmin, xmax: r1.xmax}, true
		} else {
			return xRange{xmin: r2.xmin, xmax: r2.xmax}, true
		}
	}
	return r1, false
}

func Day15() []interface{} {
	data := ReadInputAsStr(15)
	sensors := ParseSensor(data)
	return []interface{}{q15part1(sensors), q15part2(sensors)}
}

func (s Sensor) ManhattenRadius() int {
	return abs(s.x-s.bx) + abs(s.y-s.by)
}

func (s Sensor) CalculateNoBeaconRange(y int) xRange {
	// Returns the minimum and maximum x values in the row y.
	r := s.ManhattenRadius()

	// First take the y distance to the sensor
	yDist := abs(y - s.y)
	if yDist > r {
		return xRange{}
	}
	result := xRange{xmin: s.x - (r - yDist), xmax: s.x + (r - yDist)}
	return result
}

func ParseSensor(data []string) []Sensor {
	sensors := make([]Sensor, len(data))
	for idx, row := range data {
		re := regexp.MustCompile(`(-?[\d]+)`)
		submatchall := re.FindAllString(row, -1)
		x, _ := strconv.Atoi(submatchall[0])
		y, _ := strconv.Atoi(submatchall[1])
		bx, _ := strconv.Atoi(submatchall[2])
		by, _ := strconv.Atoi(submatchall[3])
		sensors[idx] = Sensor{x: x, y: y, bx: bx, by: by}
	}
	return sensors
}

func CollapseRanges(ranges []xRange) []xRange {
	solutions := []xRange{}
	resolved := []int{}
	for i, range1 := range ranges {
		for j, range2 := range ranges {
			if j <= i {
				continue
			}
			if IntInSlice(i, resolved) || IntInSlice(j, resolved) {
				continue
			}

			newRange, isCombined := range1.Combine(range2)
			if isCombined {
				solutions = append(solutions, newRange)
				resolved = append(resolved, i, j)
				break
			}
		}
		if IntInSlice(i, resolved) == false {
			solutions = append(solutions, range1)
		}
	}
	if len(solutions) < len(ranges) {
		return CollapseRanges(solutions)
	}
	return ranges
}

func TotalSpaces(ranges []xRange) int {
	total := 0
	for _, x := range ranges {
		total += x.xmax - x.xmin
	}
	return total
}

func GetRangesAtY(yValue int, sensors []Sensor) []xRange {
	ranges := []xRange{}
	for _, s := range sensors {
		xrange := s.CalculateNoBeaconRange(yValue)

		if xrange.xmin == 0 && xrange.xmax == 0 {
			continue
		}
		ranges = append(ranges, xrange)
	}
	return CollapseRanges(ranges)
}

func FindDiscontinuity(ranges []xRange, xmax int) int {
	if len(ranges) <= 1 {
		return -1
	}

	for i, range1 := range ranges {
		for j, range2 := range ranges {
			if j <= i {
				continue
			}
			if range1.xmax+2 == range2.xmin {
				return range1.xmax + 1
			}
			if range2.xmax+2 == range1.xmin {
				return range2.xmax + 1
			}
			fmt.Println(range1.xmax, range2.xmin+2)
		}
	}
	return -1
}

func q15part1(sensors []Sensor) int {
	row_num := 2000000
	ranges := GetRangesAtY(row_num, sensors)
	return TotalSpaces(ranges)
}

func TuningFrequency(x int, y int) int {
	return x*4000000 + y
}

func q15part2(sensors []Sensor) int {
	ymax := 4000000
	for y := 0; y <= ymax; y++ {
		ranges := GetRangesAtY(y, sensors)
		result := FindDiscontinuity(ranges, ymax)
		if result != -1 {
			return TuningFrequency(result, y)
		}
	}
	return -1
}
