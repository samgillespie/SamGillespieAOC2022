package answers

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Day13() []interface{} {
	data := ReadInputAsStr(13)
	return []interface{}{q13part1(data), q13part2(data)}
}

type SignalElement struct {
	value       int
	valueExists bool
	list        []*SignalElement
}

func (se SignalElement) AsString() string {
	out := ""
	if se.valueExists == true {
		out += fmt.Sprintf("%d", se.value)
	} else {
		out += "["
		for i, elem := range se.list {
			out += elem.AsString()
			if i < len(se.list)-1 {
				out += ","
			}
		}
		out += "]"
	}
	return out
}

type Comparison int

const (
	EQUAL   Comparison = 0
	GREATER            = 1
	LESS               = 2
)

func (se1 SignalElement) Compare(se2 *SignalElement) Comparison {
	if se1.valueExists {
		if se2.valueExists {
			if se1.value > se2.value {
				return GREATER
			}
			if se1.value < se2.value {
				return LESS
			}
			if se1.value == se2.value {
				return EQUAL
			}
		}
		se1_new := SignalElement{list: []*SignalElement{&se1}}
		return se1_new.Compare(se2)
	} else {
		if se2.valueExists {
			se2 = &SignalElement{list: []*SignalElement{se2}}
		}
		for i, value1 := range se1.list {

			if i >= len(se2.list) {
				return GREATER
			}
			value2 := se2.list[i]
			comp := value1.Compare(value2)
			if comp != EQUAL {
				return comp
			}
		}
		if len(se2.list) > len(se1.list) {
			return LESS
		}
	}
	return EQUAL
}

func StringOfIntsToSignalElement(input string, signalMap map[rune]*SignalElement) *SignalElement {
	// Returns "1,2,3", as []int{1,2,3}
	if input == "[]" {
		return &SignalElement{}
	}
	split := strings.Split(input[1:(len(input)-1)], ",")
	solution := make([]*SignalElement, len(split))
	for i, item := range split {
		// Check to see if in signal map
		signal, exists := signalMap[rune(item[0])]
		if exists {
			solution[i] = signal
		} else {
			value, _ := strconv.Atoi(item)
			solution[i] = &SignalElement{value: value, valueExists: true}
		}
	}
	return &SignalElement{list: solution}
}

func ParseSignalElement(input string) *SignalElement {
	// Remove
	base := &SignalElement{}
	char := 'a'

	// Extract lists, then do replacement
	regex, err := regexp.Compile(`\[([\d\w,\d\w]*)\]`)
	if err != nil {
		panic(err)
	}

	charmap := map[rune]*SignalElement{}
	for {
		if len(input) == 1 {
			index := rune(input[0])
			base = charmap[index]
			return base
		}
		if len(input) == 2 {
			return base
		}
		search := string(regex.Find([]byte(input)))
		if len(search) == 0 {
			fmt.Println("REGEX FAILED")
			fmt.Println(input)
			panic(input)
		}
		signalElement := StringOfIntsToSignalElement(string(search), charmap)
		input = strings.Replace(input, search, string(char), 1)
		charmap[char] = signalElement
		if char >= 'z' {
			char = 'A'
		} else {
			char += 1
		}
	}
}

func q13part1(data []string) int {
	solution := 0
	for i := 0; i < len(data); i += 3 {
		element_1 := ParseSignalElement(data[i])
		element_2 := ParseSignalElement(data[i+1])

		// Validate
		if element_1.AsString() != data[i] {
			fmt.Println(element_1.AsString())
			fmt.Println(data[i])
			panic("Parsing error")
		}
		if element_2.AsString() != data[i+1] {
			fmt.Println(element_2.AsString())
			fmt.Println(data[i+1])
			panic("Parsing error")
		}

		result := element_1.Compare(element_2)
		if result == 2 {
			solution += i/3 + 1
		}
	}
	return solution
}

func q13part2(data []string) int {
	bases := []*SignalElement{}
	for i := 0; i < len(data); i += 3 {
		bases = append(bases, ParseSignalElement(data[i]))
		bases = append(bases, ParseSignalElement(data[i+1]))
	}

	//Add divider packets
	bases = append(bases, ParseSignalElement("[[2]]"))
	bases = append(bases, ParseSignalElement("[[6]]"))

	// Slice magic
	sort.Slice(bases, func(i, j int) bool {
		return bases[i].Compare(bases[j]) >= 2
	})

	var index1, index2 int
	for index, row := range bases {
		if row.AsString() == "[[2]]" {
			index1 = index + 1
		}
		if row.AsString() == "[[6]]" {
			index2 = index + 1
		}
	}
	return index1 * index2
}
