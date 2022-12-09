package main

import (
	"code/aoc2021/answers"
	"flag"
	"fmt"
	"time"
)

var question int
var runProfile bool

var questionMap = map[int]func() []interface{}{
	1: answers.Day1,
	2: answers.Day2,
	3: answers.Day3,
	4: answers.Day4,
	5: answers.Day5,
	6: answers.Day6,
	7: answers.Day7,
	8: answers.Day8,
	9: answers.Day9,
}

func main() {
	parseArgs()
	if runProfile == false {
		result := SolveQuestion()
		fmt.Printf("Day %d Part 1 Answer : %v\n", question, result[0])
		fmt.Printf("Day %d Part 2 Answer : %v\n", question, result[1])
	} else {
		runs := make([]time.Duration, 0, 1000)
		for i := 0; i < 1000; i++ {
			start := time.Now()
			SolveQuestion()
			runs = append(runs, time.Since(start))
		}
		var min, max, total time.Duration
		for i, runtime := range runs {
			if i == 0 {
				min = runtime
			}
			if runtime < min {
				min = runtime
			}
			if runtime > max {
				max = runtime
			}
			total += runtime
		}
		avg := time.Duration(total.Nanoseconds() / int64(len(runs)))
		fmt.Println("min:", min, "max:", max, "avg:", avg, "total", total)
	}
}

func SolveQuestion() []interface{} {
	if question == 0 {
		times := []time.Duration{}
		for i := 1; i <= len(questionMap); i++ {
			start := time.Now()
			questionMap[i]()
			end := time.Since(start)
			times = append(times, end)
			fmt.Printf("Day %d: Time Taken %s\n", i, end)
		}
		var totalDuration time.Duration
		for _, dur := range times {
			totalDuration += dur
		}
		fmt.Printf("Total Time Taken: %s\n\n", totalDuration)
		return []interface{}{0, 0}
	} else {
		return questionMap[question]()
	}
}

func parseArgs() {
	flag.IntVar(&question, "question", 0, "Which question to answer")
	flag.BoolVar(&runProfile, "prof", false, "Whether to run a profile. If enabled runs the solution 1000 times and grabs an average, min and max runtimes")
	flag.Parse()

}
