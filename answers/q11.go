package answers

import (
	"fmt"
	"strconv"
	"strings"
)

func Day11() []interface{} {
	data := ReadInputAsStr(11)
	monkeys := ParseMonkey(data)
	monkeys2 := make([]Monkey, len(monkeys))
	copy(monkeys2, monkeys) // Return to original state
	return []interface{}{q11part1(monkeys), q11part2(monkeys2)}
}

type MonkeyOperation struct {
	add      bool
	multiply bool
	double   bool
	value    int
}

func (mo MonkeyOperation) DoOperation(value int) int {
	if mo.add == true {
		return value + mo.value
	} else if mo.multiply == true {
		return value * mo.value
	} else if mo.double == true {
		return value * value
	}
	panic("Should not happen")
}

type Monkey struct {
	num         int
	items       []int
	operation   MonkeyOperation
	test        int // divisible by
	trueMonkey  int
	falseMonkey int
	counter     int
}

// Returns item number and target monkey
func (m *Monkey) InspectItem(relief_lcm int) (int, int) {
	m.counter += 1
	var item int
	item, m.items = m.items[0], m.items[1:] // Remove from queue
	item = m.operation.DoOperation(item)
	if relief_lcm == 0 {
		item = item / 3 // Relief
	} else {
		item = item % relief_lcm
	}

	if item%m.test == 0 {
		return item, m.trueMonkey
	} else {
		return item, m.falseMonkey
	}
}

func CalculateMonkeyBusiness(monkeys []Monkey) int {
	counters := []int{}
	for _, monkey := range monkeys {
		counters = append(counters, monkey.counter)
	}

	index, monkeyBusiness1 := maxSlice(counters)
	counters[index] = 0
	_, monkeyBusiness2 := maxSlice(counters)
	return monkeyBusiness1 * monkeyBusiness2
}

func ParseMonkey(data []string) []Monkey {
	monkeys := []Monkey{} //Hey Hey
	activeMonkey := Monkey{}
	for _, row := range data {
		// Clean input
		cleaned := strings.Trim(row, " ")
		cleaned = strings.ReplaceAll(cleaned, ":", "")
		cleaned = strings.ReplaceAll(cleaned, ",", "")
		splitted := strings.Split(cleaned, " ")
		switch splitted[0] {
		case "Monkey":
			num, _ := strconv.Atoi(splitted[1])
			activeMonkey.num = num
		case "Starting":
			items := []int{}
			for _, item_num := range splitted[2:] {
				item, _ := strconv.Atoi(item_num)
				items = append(items, item)
			}
			activeMonkey.items = items
		case "Operation":
			operation := MonkeyOperation{}
			if splitted[5] == "old" {
				operation.double = true
			} else {
				if splitted[4] == "*" {
					operation.multiply = true
				} else if splitted[4] == "+" {
					operation.add = true
				} else {
					panic(splitted)
				}
				value, _ := strconv.Atoi(splitted[5])
				operation.value = value
			}
			activeMonkey.operation = operation
		case "Test":
			test, _ := strconv.Atoi(splitted[3])
			activeMonkey.test = test
		case "If":
			if splitted[1] == "true" {
				trueMonkey, _ := strconv.Atoi(splitted[5])
				activeMonkey.trueMonkey = trueMonkey
			} else if splitted[1] == "false" {
				falseMonkey, _ := strconv.Atoi(splitted[5])
				activeMonkey.falseMonkey = falseMonkey
			}
		case "":
			monkeys = append(monkeys, activeMonkey)
			activeMonkey = Monkey{}
		default:
			panic(splitted[0])
		}
	}
	monkeys = append(monkeys, activeMonkey)
	return monkeys
}

func q11part1(monkeys []Monkey) int {
	for round_num := 0; round_num < 20; round_num++ {
		for monkey_num, monkey := range monkeys {
			for {
				if len(monkey.items) == 0 {
					monkeys[monkey_num] = monkey
					break
				}
				item, target := monkey.InspectItem(0)
				monkeys[target].items = append(monkeys[target].items, item)
			}
		}
	}
	return CalculateMonkeyBusiness(monkeys)
}

func q11part2(monkeys []Monkey) int {
	lowestCommonMultiple := 1
	for _, monkey := range monkeys {
		lowestCommonMultiple *= monkey.test
	}
	for round_num := 0; round_num < 10000; round_num++ {
		for monkey_num, monkey := range monkeys {
			for {
				if len(monkey.items) == 0 {
					monkeys[monkey_num] = monkey
					break
				}
				item, target := monkey.InspectItem(lowestCommonMultiple)
				monkeys[target].items = append(monkeys[target].items, item)
			}
		}
	}
	for _, monkey := range monkeys {
		fmt.Println(monkey.num, monkey.counter)
	}
	return CalculateMonkeyBusiness(monkeys)
}

// Ok so the worry levels need to be constrained to the t
