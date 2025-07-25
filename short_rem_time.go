// Shorted Remaining Time First Algorithm
// Process with least remaining burst time is given priority
// But if another process arrives with lower burst time, then curr process
// is paused to execute the newly arrived process with lower burst time

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func shortestRemainingTimeFirst(processes []Process) {
	n := len(processes)
	waiting_time := make([]float64, n)
	turn_around_time := make([]float64, n)
	completion_time := make([]float64, n)
	remaining_burst := make([]float64, n)

	sortByArrival := func(a, b Process) int {
		return cmp.Compare(a.arrival, b.arrival)
	}
	slices.SortFunc(processes, sortByArrival)

	for i := range n {
		remaining_burst[i] = processes[i].burst
	}

	ready_queue := []int{}
	curr_time := 0.0
	last_added := -1
	total_tat_time := 0.0
	total_wait_time := 0.0

	for last_added < n-1 || len(ready_queue) > 0 {
		// Add all the processes arrived to ready queue (if not addded)
		i := last_added + 1
		for ; i < n; i++ {
			if processes[i].arrival > curr_time {
				break
			}
			ready_queue = append(ready_queue, i)
		}
		last_added = i - 1

		if len(ready_queue) == 0 && last_added < n-1 {
			ready_queue = append(ready_queue, last_added+1)
			curr_time = processes[last_added+1].arrival
			var j int
			for j = last_added + 2; j < n; j++ {
				if processes[j].arrival > curr_time {
					break
				}
				ready_queue = append(ready_queue, j)
			}
			last_added = j - 1
		}

		min_burst_idx := 0

		// Find process with least remaining burst time
		for j := range len(ready_queue) {
			idx := ready_queue[j]
			if remaining_burst[idx] < remaining_burst[ready_queue[min_burst_idx]] {
				min_burst_idx = j
			}
		}
		min_burst := ready_queue[min_burst_idx]

		// Execute process with shortest remaining time (at index `idx`)
		remaining_burst[min_burst]--
		curr_time++

		if remaining_burst[min_burst] == 0 {
			completion_time[min_burst] = curr_time
			turn_around_time[min_burst] = completion_time[min_burst] - processes[min_burst].arrival
			waiting_time[min_burst] = turn_around_time[min_burst] - processes[min_burst].burst

			ready_queue = append(ready_queue[:min_burst_idx], ready_queue[min_burst_idx+1:]...)
			total_tat_time += turn_around_time[min_burst]
			total_wait_time += waiting_time[min_burst]
		}
	}

	fmt.Println(remaining_burst)

	avg_tat_time := total_tat_time / float64(n)
	avg_wait_time := total_wait_time / float64(n)

	printAlgoResult("Shortest Remaining Time First Algorithm", avg_wait_time, avg_tat_time)
}
