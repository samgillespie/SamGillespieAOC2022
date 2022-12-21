package answers

import (
	"fmt"
	"strconv"
	"strings"
)

func Day21() []interface{} {
	data := ReadInputAsStr(21)
	return []interface{}{q21part1(data), q21part2(data)}
}

func ParseScreamingMonkeys(data []string) map[string]*ScreamingMonkey {
	monkey_map := map[string]*ScreamingMonkey{}
	for _, row := range data {
		split := strings.Split(row, ":")

		value, err := strconv.Atoi(strings.Trim(split[1], " "))
		monkey := ScreamingMonkey{name: strings.Trim(split[0], " ")}
		if err != nil {
			split2 := strings.Split(strings.Trim(split[1], " "), " ")
			monkey.monkey_a_str = split2[0]
			monkey.monkey_b_str = split2[2]
			monkey.operator = split2[1]
		} else {
			monkey.value = value
		}
		monkey_map[monkey.name] = &monkey
	}
	for _, monkey := range monkey_map {
		if monkey.monkey_a_str != "" {
			monkey.monkey_a, _ = monkey_map[monkey.monkey_a_str]
			monkey.monkey_b, _ = monkey_map[monkey.monkey_b_str]
		}
	}
	return monkey_map
}

type ScreamingMonkey struct {
	name         string
	monkey_a     *ScreamingMonkey
	monkey_b     *ScreamingMonkey
	monkey_a_str string
	monkey_b_str string
	value        int
	operator     string
}

func (sm ScreamingMonkey) Scream() int {
	if sm.value != 0 {
		return sm.value
	}
	switch sm.operator {

	case "-":
		return sm.monkey_a.Scream() - sm.monkey_b.Scream()
	case "+":
		return sm.monkey_a.Scream() + sm.monkey_b.Scream()
	case "*":
		return sm.monkey_a.Scream() * sm.monkey_b.Scream()
	case "/":
		return sm.monkey_a.Scream() / sm.monkey_b.Scream()
	case "=":
		a := sm.monkey_a.Scream()
		b := sm.monkey_b.Scream()
		fmt.Printf("%d, %d\n\n", a, b)
		if a == b {
			return 1
		}
		return 0
	}
	return -1
}

func q21part1(data []string) int {
	monkey := ParseScreamingMonkeys(data)
	return monkey["root"].Scream()
}

func q21part2(data []string) int {
	monkey := ParseScreamingMonkeys(data)
	monkey["root"].operator = "="
	human := monkey["humn"]
	stepsize := 100000
	human.value = 0
	monkey_previous := monkey["root"].monkey_a.Scream() // Assume only monkey_a changes
	target_value := monkey["root"].monkey_b.Scream()
	// Some kind of euler's method
	for stepsize != 0 {
		human.value += stepsize
		monkey_next := monkey["root"].monkey_a.Scream()
		difference := monkey_next - monkey_previous
		rate_of_change := difference / stepsize

		// Difference = how many does the number change when you change stepsize
		target_change := target_value - monkey_next
		stepsize = target_change / rate_of_change
		fmt.Println(human.value, stepsize, rate_of_change, target_change)
		monkey_previous = monkey_next
	}
	fmt.Println(human.value, monkey["root"].monkey_a.Scream(), monkey["root"].monkey_b.Scream(), monkey["root"].Scream())
	return human.value
}
