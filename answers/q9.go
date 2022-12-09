package answers

import (
	"fmt"
	"strconv"
	"strings"
)

type RopeInstruction struct {
	Direction string
	Distance  int
}

func Day9() []interface{} {
	data := ReadInputAsStr(9)
	instructions := []RopeInstruction{}
	for _, rope := range data {
		split := strings.Split(rope, " ")
		dist, _ := strconv.Atoi(split[1])
		instructions = append(instructions, RopeInstruction{Direction: split[0], Distance: dist})
	}
	return []interface{}{q9part1(instructions), q9part2(instructions)}
}

type Rope struct {
	x      int
	y      int
	parent *Rope
}

func (r Rope) Print(idx string) {
	fmt.Printf("idx: %s, x:%d, y:%d\n", idx, r.x, r.y)
}

func (r1 Rope) Difference(r2 Rope) Vector {
	return Vector{x: r1.x - r2.x, y: r1.y - r2.y}
}

func (r Rope) MoveHead(direction string) Rope {
	if direction == "D" {
		r.y -= 1
	}
	if direction == "U" {
		r.y += 1
	}
	if direction == "R" {
		r.x += 1
	}
	if direction == "L" {
		r.x -= 1
	}
	return r
}

func (r Rope) FollowParent() Rope {
	distance := r.parent.Difference(r)
	toMoveX := 0
	toMoveY := 0
	if distance.x > 1 {
		toMoveX = 1
		if distance.y > 0 {
			toMoveY = 1
		}
		if distance.y < 0 {
			toMoveY = -1
		}
	}
	if distance.x < -1 {
		toMoveX = -1
		if distance.y > 0 {
			toMoveY = 1
		}
		if distance.y < 0 {
			toMoveY = -1
		}
	}

	if distance.y > 1 {
		toMoveY = 1
		if distance.x > 0 {
			toMoveX = 1
		}
		if distance.x < 0 {
			toMoveX = -1
		}
	}
	if distance.y < -1 {
		toMoveY = -1
		if distance.x > 0 {
			toMoveX = 1
		}
		if distance.x < 0 {
			toMoveX = -1
		}
	}
	r.x += toMoveX
	r.y += toMoveY
	return r
}

func q9part1(instructions []RopeInstruction) int {
	head := Rope{x: 0, y: 0}
	tail := Rope{x: 0, y: 0, parent: &head}
	visited := map[Rope]bool{}
	for _, instruction := range instructions {
		for i := 0; i < instruction.Distance; i++ {
			head = head.MoveHead(instruction.Direction)
			tail = tail.FollowParent()
			visited[tail] = true
		}
	}
	return len(visited)
}

func q9part2(instructions []RopeInstruction) int {
	head := Rope{x: 0, y: 0}
	tails := []Rope{}
	for i := 0; i < 9; i++ {
		tails = append(tails, Rope{})
	}
	tails[0].parent = &head
	// Janky to get the pointers to work
	for i := 1; i < 9; i++ {
		tail := tails[i]
		tail.parent = &tails[i-1]
		tails[i] = tail
	}

	visited := map[Rope]bool{}
	for _, instruction := range instructions {
		for i := 0; i < instruction.Distance; i++ {
			head = head.MoveHead(instruction.Direction)
			for idx := range tails {
				tail := tails[idx].FollowParent()
				tails[idx] = tail
			}
			visited[tails[8]] = true
		}
	}
	return len(visited)
}
