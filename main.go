package main

import (
	"fmt"
	"strings"
)

type Process struct {
	arrival  float64
	burst    float64
	priority int
}

func main() {

	var processes []Process = []Process{
		{0, 1, 2}, {1, 7, 3}, {2, 22, 12}, {3, 3, 5}, {4, 9, 11}, {5, 15, 9}, {6, 26, 7}, {7, 17, 6}, {8, 8, 1}, {9, 19, 8},
	}

	firstComeFirstServe(processes)
	shortestJobFirst(processes)
	longestJobFirst(processes)
	highestResponseRatioNext(processes)
	priorityScheduling(processes)
	nonPreemptivePriorityScheduling(processes)

}

func printAlgoResult(algo_name string, avg_wait_time, avg_tat_time float64) {
	fmt.Println()
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println(algo_name)
	fmt.Println("Average waiting time: ", avg_wait_time)
	fmt.Println("Average turn around time: ", avg_tat_time)
	fmt.Println(strings.Repeat("-", 40))
	fmt.Println()
}
