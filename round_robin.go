// Round Robin algorithm implementation
// We will have 2 queues ready and running

package main

import (
	"cmp"
	"slices"
)

func roundRobin(processes []Process, quantum float64) {
	n := len(processes)
	waiting_time := make([]float64, n)
	turn_around_time := make([]float64, n)
	completion_time := make([]float64, n)
	burst_times := make([]float64, n)

	sortByArrival := func(a, b Process) int {
		return cmp.Compare(a.arrival, b.arrival)
	}

	slices.SortFunc(processes, sortByArrival)

	for i := 0; i < n; i++ {
		burst_times[i] = processes[i].burst
	}

	var ready_queue []int = []int{0}

	last_added := 0
	curr_time := processes[0].arrival

	total_tat_time := 0.0
	total_wait_time := 0.0

	for len(ready_queue) > 0 {
		// Add all the available processes to ready queue

		idx := ready_queue[0]

		if burst_times[idx] <= quantum {

			curr_time += burst_times[idx]
			burst_times[idx] = 0

			completion_time[idx] = curr_time
			turn_around_time[idx] = completion_time[idx] - processes[idx].arrival
			waiting_time[idx] = turn_around_time[idx] - processes[idx].burst

			total_wait_time += waiting_time[idx]
			total_tat_time += turn_around_time[idx]

		} else {
			burst_times[idx] -= quantum
			curr_time += quantum
		}

		// remove current process
		ready_queue = ready_queue[1:]

		// Add all the available processes
		i := last_added + 1
		for i < n {
			if processes[i].arrival > curr_time {
				break
			}
			ready_queue = append(ready_queue, i)
			i += 1
		}
		last_added = i - 1

		// If current process did not completed add it again at the last of queue
		if burst_times[idx] > 0 {
			ready_queue = append(ready_queue, idx)
		}

		if len(ready_queue) == 0 && last_added < n-1 {
			curr_time = processes[last_added+1].arrival
			ready_queue = append(ready_queue, last_added+1)
			last_added++
		}
	}

	avg_wait_time := total_wait_time / float64(n)
	avg_tat_time := total_tat_time / float64(n)
	printAlgoResult("Round Robin Algorithm", avg_wait_time, avg_tat_time)
}
