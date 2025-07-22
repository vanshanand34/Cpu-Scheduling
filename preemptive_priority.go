// Priority Scheduling algorithm implementation
// process with highest priority gets executed first
// processes with same priority gets exectuted in the order of their arrival
// on a First Come First Server basis

// Lower number = Higher Priority
package main

import (
	"cmp"
	"math"
	"slices"
)

func nonPreemptivePriorityScheduling(processes []Process) {
	var n int = len(processes)

	waiting_time := make([]float64, n)
	completion_time := make([]float64, n)
	turn_around_time := make([]float64, n)

	sortByPriority := func(a, b Process) int {
		return cmp.Compare(a.priority, b.priority)
	}

	slices.SortFunc(processes, sortByPriority)

	curr_time := 0.0
	total_tat_time := 0.0
	total_wait_time := 0.0
	total_burst_time := 0.0

	var remaining_burst = make([]float64, n)

	for i := range n {
		total_burst_time += processes[i].burst
		remaining_burst[i] = processes[i].burst
	}

	for i := range int(math.Ceil(total_burst_time)) {
		// select process with least priority available at this time (curr_time)
		idx := i
		idx = -1

		for j := range n {
			if remaining_burst[j] == 0 {
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
				if remaining_burst[j] > 0 {
					idx = j
					break
				}
			}
		}

		remaining_burst[idx]--
		curr_time++

		if remaining_burst[idx] == 0 {
			// process completed
			completion_time[idx] = curr_time
			turn_around_time[idx] = completion_time[idx] - processes[idx].arrival
			waiting_time[idx] = turn_around_time[idx] - processes[idx].burst

			total_tat_time += turn_around_time[idx]
			total_wait_time += waiting_time[idx]
		}

	}

	avg_wait_time := total_wait_time / float64(n)
	avg_tat_time := total_tat_time / float64(n)

	printAlgoResult("Preemptive Priority Scheduling Algorithm", avg_wait_time, avg_tat_time)
}
