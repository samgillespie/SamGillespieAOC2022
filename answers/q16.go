package answers

import (
	"fmt"
	"strconv"
	"strings"
)

// Only link valves with flow rates together
type ValveLink struct {
	origin   *Valve
	target   *Valve
	distance int
}

type Valve struct {
	name      string
	rate      int
	exits_str []string // Used initially, then discarded
	exits     []ValveLink
}

type ValvePath struct {
	valvesOpened  []*Valve
	position      *Valve
	totalPressure int
	timeRemaining int
}

func PrintValveLink(links []ValveLink) {
	for _, link := range links {
		fmt.Printf("Valve: %s to Valve: %s has Distance %d\n", link.origin.name, link.target.name, link.distance)
	}
}

func Day16() []interface{} {
	data := ReadInputAsStr(16)
	valves := ParseValves(data)
	return []interface{}{q16part1(valves), q16part2(valves)}
}

func BreadthStep(path ValvePath) []ValvePath {
	solution := []ValvePath{}
	for _, valveLink := range path.position.exits {
		if ValveInSlice(valveLink.target, path.valvesOpened) {
			continue
		}
		timeRemaining := path.timeRemaining - valveLink.distance
		if timeRemaining < 0 {
			continue
		}

		newValvesOpened := make([]*Valve, len(path.valvesOpened))
		copy(newValvesOpened, path.valvesOpened)
		newValvesOpened = append(newValvesOpened, valveLink.target)

		addedPressure := valveLink.target.rate * timeRemaining

		newPath := ValvePath{
			valvesOpened:  newValvesOpened,
			position:      valveLink.target,
			totalPressure: path.totalPressure + addedPressure,
			timeRemaining: timeRemaining,
		}
		solution = append(solution, newPath)
	}
	return solution
}

func ValveInSlice(valve *Valve, slice []*Valve) bool {
	for _, check := range slice {
		if check.name == valve.name {
			return true
		}
	}
	return false
}

func q16part1(valves map[string]*Valve) int {
	best := 0
	startPos := valves["AA"]
	start := ValvePath{position: startPos, timeRemaining: 30}
	currentStep := []ValvePath{start}
	for len(currentStep) > 0 {
		nextStep := []ValvePath{}
		for _, step := range currentStep {
			next := BreadthStep(step)
			if len(next) == 0 {
				if step.totalPressure > best {
					best = step.totalPressure
				}
			} else {
				nextStep = append(nextStep, next...)
			}
		}
		currentStep = nextStep
	}
	return best
}

func MutuallyExclusive(list1 []*Valve, list2 []*Valve) bool {
	for _, elem1 := range list1 {
		for _, elem2 := range list2 {
			if elem1.name == elem2.name {
				return false
			}
		}
	}
	return true
}

func q16part2(valves map[string]*Valve) int {
	startPos := valves["AA"]
	start := ValvePath{position: startPos, timeRemaining: 26}
	currentStep := []ValvePath{start}
	allSolutions := []ValvePath{}
	// Calculate all paths
	for len(currentStep) > 0 {
		nextStep := []ValvePath{}
		for _, step := range currentStep {
			next := BreadthStep(step)
			if len(next) == 0 {
				allSolutions = append(allSolutions, step)
			} else {
				nextStep = append(nextStep, next...)
			}
		}
		currentStep = nextStep
	}

	// Now find all possible combinations of two paths
	best := 0
	for i, step1 := range allSolutions {
		for j, step2 := range allSolutions {
			if j >= i {
				continue //Upper triange
			}
			if MutuallyExclusive(step1.valvesOpened, step2.valvesOpened) == true {
				totalPressure := step1.totalPressure + step2.totalPressure
				if totalPressure > best {
					best = totalPressure
				}
			}
		}
	}
	return best
}

func ParseValves(data []string) map[string]*Valve {
	mapValves := map[string]*Valve{}
	for _, row := range data {
		filter := strings.Split(row, "=")
		name := strings.Split(filter[0], " ")[1]

		filter2 := strings.Split(filter[1], ";")
		rate, _ := strconv.Atoi(filter2[0])

		splitTunnels := strings.Split(filter2[1], " ")[5:]
		exits_str := []string{}
		for _, tunnel := range splitTunnels {
			exits_str = append(exits_str, strings.Replace(tunnel, ",", "", 1))
		}
		valve := Valve{name: name, rate: rate, exits_str: exits_str}
		mapValves[name] = &valve
	}

	//Calculate ValveLinks for any rate
	for valveName, valve := range mapValves {
		if valve.rate == 0 && valve.name != "AA" {
			continue
		}
		steps := 1
		exits := valve.exits_str
		distances := map[*Valve]int{valve: 0}
		for len(distances) < len(mapValves) {
			nextStep := []string{}
			for _, nextValveStr := range exits {
				nextValve, ok := mapValves[nextValveStr]
				if !ok {
					panic("wat")
				}
				_, inMap := distances[nextValve]
				if inMap == false {
					distances[nextValve] = steps
				}
				nextStep = append(nextStep, nextValve.exits_str...)
			}
			exits = nextStep
			steps += 1
		}
		// Convert distances into links
		links := []ValveLink{}
		for targetValve, distance := range distances {
			if targetValve.rate == 0 {
				// Can never target AA
				continue
			}
			links = append(links, ValveLink{
				origin:   valve,
				target:   targetValve,
				distance: distance + 1, // +1 for actually opening the valve
			})
		}
		valve.exits = links
		mapValves[valveName] = valve
	}

	// Only return the valves with rate > 0.
	for valveName, valve := range mapValves {
		if valve.rate == 0 && valve.name != "AA" {
			delete(mapValves, valveName)
		}
	}
	return mapValves
}
