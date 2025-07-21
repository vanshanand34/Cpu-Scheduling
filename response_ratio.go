// Program to implement HRRN (Highest Response Ratio Next) Algorithm
package main

import (
	"cmp"
	"slices"
)

func highestResponseRatioNext(processes []Process) {
	n := len(processes)

	// sorting processes by increasing arrival time
	sortByArrival := func(a, b Process) int {
		return cmp.Compare(a.arrival, b.arrival)
	}
	slices.SortFunc(processes, sortByArrival)

	response_ratios := make([]float64, n)
	turn_around_time := make([]float64, n)
	waiting_time := make([]float64, n)
	completion_time := make([]float64, n)
	visited := make([]bool, n)

	waiting_time[0] = 0
	turn_around_time[0] = processes[0].burst
	completion_time[0] = processes[0].burst + processes[0].arrival
	response_ratios[0] = 0

	curr_time := completion_time[0]
	visited[0] = true

	total_wait_time := 0.0
	total_tat_time := 0.0

	for i := 1; i < n; i++ {
		max_ratio := -1.0
		idx := -1
		for i := range n {

			if visited[i] {
				continue
			}
			// determining response ratios of all remaining processes
			wait_time := max(curr_time-processes[i].arrival, 0)
			response_ratios[i] = (wait_time + processes[i].burst) / processes[i].burst
			if response_ratios[i] > max_ratio {
				max_ratio = response_ratios[i]
				idx = i
			}
		}

		// execute process with max response ratio
		curr_process := processes[idx]
		waiting_time[idx] = max(curr_time-curr_process.arrival, 0)
		turn_around_time[idx] = waiting_time[idx] + curr_process.burst
		completion_time[idx] = curr_process.arrival + turn_around_time[idx]

		// mark as visited
		visited[idx] = true

		// update curr time
		curr_time = completion_time[idx]

		total_tat_time += turn_around_time[idx]
		total_wait_time += waiting_time[idx]
	}

	avg_wait_time := total_wait_time / float64(n)
	avg_tat_time := total_tat_time / float64(n)

	printAlgoResult("Highest Response Ratio Algorithm", avg_wait_time, avg_tat_time)
}

// Response Ratio = (W + S)/S
// W : Waiting time of the process so far
// S : Burst time of the process.
