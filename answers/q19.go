package answers

import (
	"fmt"
	"strconv"
	"strings"
)

func Day19() []interface{} {
	data := ReadInputAsStr(19)
	blueprints := ParseBlueprints(data)
	return []interface{}{q19part1(blueprints), q19part2(blueprints)}
}

func ParseBlueprints(data []string) []Blueprint {
	blueprints := []Blueprint{}
	number := 1
	for _, row := range data {
		step1 := strings.Split(row, ":")[1]
		step2 := strings.Split(step1, ".")
		oreRobotore, _ := strconv.Atoi(strings.Split(step2[0], " ")[5])
		oreRobotCost := Cost{Ore: oreRobotore}

		clayRobotore, _ := strconv.Atoi(strings.Split(step2[1], " ")[5])
		clayRobotCost := Cost{Ore: clayRobotore}

		obsidianRobotore, _ := strconv.Atoi(strings.Split(step2[2], " ")[5])
		obsidianRobotclay, _ := strconv.Atoi(strings.Split(step2[2], " ")[8])
		obsidianRobotCost := Cost{Ore: obsidianRobotore, Clay: obsidianRobotclay}

		geodeRobotore, _ := strconv.Atoi(strings.Split(step2[3], " ")[5])
		geodeRobotobsidian, _ := strconv.Atoi(strings.Split(step2[3], " ")[8])
		geodeRobotCost := Cost{Ore: geodeRobotore, Obsidian: geodeRobotobsidian}
		blueprint := Blueprint{
			Number:        number,
			OreRobot:      oreRobotCost,
			ClayRobot:     clayRobotCost,
			ObsidianRobot: obsidianRobotCost,
			GeodeRobot:    geodeRobotCost,
		}
		blueprint.CalculateMax()
		blueprints = append(blueprints, blueprint)
		number++
	}
	return blueprints
}

type Cost struct {
	Ore      int
	Clay     int
	Obsidian int
}

type Blueprint struct {
	Number        int
	OreRobot      Cost
	ClayRobot     Cost
	ObsidianRobot Cost
	GeodeRobot    Cost

	// The max we'll ever need.  If the most expensive robot costs 3 Ore, then getting more than 3 is useless
	MaxOre      int
	MaxClay     int
	MaxObsidian int
}

type CurrentStep struct {
	Blueprint *Blueprint
	Time      int
	//Robots
	OreRobot      int
	ClayRobot     int
	ObsidianRobot int
	GeodeRobot    int

	//Resources
	Ore      int
	Clay     int
	Obsidian int
	Geodes   int
}

func (cs CurrentStep) Copy() CurrentStep {
	return CurrentStep{
		Blueprint:     cs.Blueprint,
		Time:          cs.Time,
		OreRobot:      cs.OreRobot,
		ClayRobot:     cs.ClayRobot,
		ObsidianRobot: cs.ObsidianRobot,
		GeodeRobot:    cs.GeodeRobot,
		Ore:           cs.Ore,
		Clay:          cs.Clay,
		Obsidian:      cs.Obsidian,
		Geodes:        cs.Geodes,
	}
}

func (cs CurrentStep) Print() {
	fmt.Printf("Blueprint: %d, Time: %d, OreRobot: %d, ClayRobot: %d, ObsidianRobot: %d GeodeRobot: %d, Ore: %d, Clay: %d, Obsidian: %d, Geodes: %d\n",
		cs.Blueprint.Number,
		cs.Time,
		cs.OreRobot,
		cs.ClayRobot,
		cs.ObsidianRobot,
		cs.GeodeRobot,
		cs.Ore,
		cs.Clay,
		cs.Obsidian,
		cs.Geodes,
	)
}

func Intmax(a int, b int, c int, d int) int {
	// dumb
	if a >= b && a >= c && a >= d {
		return a
	} else if b >= c && b >= d {
		return b
	} else if c >= d {
		return c
	}
	return d
}

func (b *Blueprint) CalculateMax() {
	b.MaxOre = Intmax(b.OreRobot.Ore, b.ClayRobot.Ore, b.ObsidianRobot.Ore, b.GeodeRobot.Ore)
	b.MaxClay = Intmax(b.OreRobot.Clay, b.ClayRobot.Clay, b.ObsidianRobot.Clay, b.GeodeRobot.Clay)
	b.MaxObsidian = Intmax(b.OreRobot.Obsidian, b.ClayRobot.Obsidian, b.ObsidianRobot.Obsidian, b.GeodeRobot.Obsidian)
}

func (cs CurrentStep) CanBuy(cost Cost) bool {
	return cs.Ore >= cost.Ore && cs.Clay >= cost.Clay && cs.Obsidian >= cost.Obsidian
}

func (cs CurrentStep) SpendResourcesEqualToCost(cost Cost) CurrentStep {
	csNew := cs.Copy()
	csNew.Ore -= cost.Ore
	csNew.Clay -= cost.Clay
	csNew.Obsidian -= cost.Obsidian
	return csNew
}

func (cs *CurrentStep) MineResources() {
	cs.Ore += cs.OreRobot
	cs.Clay += cs.ClayRobot
	cs.Obsidian += cs.ObsidianRobot
	cs.Geodes += cs.GeodeRobot
}

//Depth first search
func (cs *CurrentStep) DFSStep(maxTime int) CurrentStep {
	nextSteps := []CurrentStep{}

	bestGeodes := BESTGEODES[cs.Time]
	// If another paths has more geodes, abandon this fork
	if cs.Geodes < bestGeodes {
		return *cs
	}
	if cs.Geodes > bestGeodes {
		BESTGEODES[cs.Time] = cs.Geodes
	}
	cs.Time++
	if cs.Time > maxTime {
		return *cs
	}

	if cs.CanBuy(cs.Blueprint.GeodeRobot) {
		newState := cs.SpendResourcesEqualToCost(cs.Blueprint.GeodeRobot)
		newState.GeodeRobot++
		newState.Geodes-- // Robot isn't ready till next iteration, so just minus 1 for now.
		newState.MineResources()
		nextSteps = append(nextSteps, newState.DFSStep(maxTime))
	} else {
		if cs.CanBuy(cs.Blueprint.ObsidianRobot) && cs.ObsidianRobot < cs.Blueprint.MaxObsidian {
			newState := cs.SpendResourcesEqualToCost(cs.Blueprint.ObsidianRobot)
			newState.ObsidianRobot++
			newState.Obsidian-- // Robot isn't ready till next iteration, so just minus 1 for now.
			newState.MineResources()
			nextSteps = append(nextSteps, newState.DFSStep(maxTime))
		}
		if cs.CanBuy(cs.Blueprint.ClayRobot) && cs.ClayRobot < cs.Blueprint.MaxClay {
			newState := cs.SpendResourcesEqualToCost(cs.Blueprint.ClayRobot)
			newState.ClayRobot++
			newState.Clay-- // Robot isn't ready till next iteration, so just minus 1 for now.
			newState.MineResources()
			nextSteps = append(nextSteps, newState.DFSStep(maxTime))
		}
		if cs.CanBuy(cs.Blueprint.OreRobot) && cs.OreRobot < cs.Blueprint.MaxOre {
			newState := cs.SpendResourcesEqualToCost(cs.Blueprint.OreRobot)
			newState.OreRobot++
			newState.Ore-- // Robot isn't ready till next iteration, so just minus 1 for now.
			newState.MineResources()
			nextSteps = append(nextSteps, newState.DFSStep(maxTime))
		}
		noPurchase := cs.Copy()
		noPurchase.MineResources()
		nextSteps = append(nextSteps, noPurchase.DFSStep(maxTime))
	}

	_, maxStep := MaxGeodes(nextSteps)
	return maxStep
}

func MaxGeodes(steps []CurrentStep) (int, CurrentStep) {
	maxGeodes := 0
	var maxStep CurrentStep
	for _, step := range steps {
		if step.Geodes > maxGeodes {
			maxGeodes = step.Geodes
			maxStep = step
		}
	}
	return maxGeodes, maxStep
}

var BESTGEODES map[int]int

func SimulateBlueprint(blueprint Blueprint, maxStep int) int {
	step := CurrentStep{
		Blueprint: &blueprint,
		OreRobot:  1,
	}
	// Keep Track of the best geodes
	BESTGEODES = map[int]int{}
	for i := 0; i < maxStep; i++ {
		BESTGEODES[i] = 0
	}
	solution := step.DFSStep(maxStep)
	return solution.Geodes
}

func q19part1(blueprints []Blueprint) int {
	qualityScore := 0
	for _, blueprint := range blueprints {
		maxGeodes := SimulateBlueprint(blueprint, 24)
		qualityScore += maxGeodes * blueprint.Number
		fmt.Println(blueprint.Number, maxGeodes*blueprint.Number)
	}
	return qualityScore
}

func q19part2(blueprints []Blueprint) int {
	solution := 1
	for i, blueprint := range blueprints {
		if i >= 3 {
			break
		}
		maxGeodes := SimulateBlueprint(blueprint, 32)
		solution *= maxGeodes
		fmt.Println(blueprint.Number, maxGeodes)
	}
	return solution
}
