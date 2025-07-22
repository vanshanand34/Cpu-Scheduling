// Priority Scheduling algorithm implementation
// process with highest priority gets executed first
// processes with same priority gets exectuted in the order of their arrival
// on a First Come First Server basis

// Lower number = Higher Priority
package main

import (
	"cmp"
	"fmt"
	"slices"
)

func priorityScheduling(processes []Process) {
	var n int = len(processes)
	fmt.Print()
	waiting_time := make([]float64, n)
	completion_time := make([]float64, n)
	turn_around_time := make([]float64, n)
	executed := make([]bool, n)

	sortByPriority := func(a, b Process) int {
		return cmp.Compare(a.priority, b.priority)
	}

	slices.SortFunc(processes, sortByPriority)

	curr_time := 0.0
	total_tat_time := 0.0
	total_wait_time := 0.0

	for i := range n {
		// select process with least priority available at this time (curr_time)
		idx := i
		idx = -1

		for j := range n {
			if executed[j] {
				continue
			}

			wait_time := processes[j].arrival - curr_time
			if wait_time <= 0 {
				idx = j
				break
			}
		}

		if idx == -1 {
			// no process available at curr time
			// choose first unexecuted process
			for j := range n {
				if !executed[j] {
					idx = j
					break
				}
			}
		}

		completion_time[idx] = max(curr_time, processes[idx].arrival) + processes[idx].burst
		turn_around_time[idx] = completion_time[idx] - processes[idx].arrival
		waiting_time[idx] = turn_around_time[idx] - processes[idx].burst
		executed[idx] = true
		curr_time = completion_time[idx]

		total_tat_time += turn_around_time[idx]
		total_wait_time += waiting_time[idx]
	}

	avg_wait_time := total_wait_time / float64(n)
	avg_tat_time := total_tat_time / float64(n)

	printAlgoResult("Non Preemptive Priority Scheduling Algorithm", avg_wait_time, avg_tat_time)
}
