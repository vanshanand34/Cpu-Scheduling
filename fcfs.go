package main

import (
	"cmp"
	"slices"
)

func firstComeFirstServe(processes []Process) {
	sortByArrival := func(a, b Process) int { return cmp.Compare(a.arrival, b.arrival) }
	slices.SortFunc(processes, sortByArrival)

	var n int = len(processes)
	completion_time := make([]float64, n)
	waiting_time := make([]float64, n)
	turn_around_time := make([]float64, n)

	completion_time[0] = processes[0].burst
	turn_around_time[0] = processes[0].burst

	var total_wait_time float64
	var total_tat_time float64 = turn_around_time[0]

	for i := 1; i < n; i++ {
		completion_time[i] = max(completion_time[i-1], processes[i].arrival) + processes[i].burst
		turn_around_time[i] = completion_time[i] - processes[i].arrival
		waiting_time[i] = turn_around_time[i] - processes[i].burst

		total_tat_time += turn_around_time[i]
		total_wait_time += waiting_time[i]
	}

	avg_waiting_time := total_wait_time / float64(n)
	avg_turn_around_time := total_tat_time / float64(n)

	printAlgoResult("First Come First Serve", avg_waiting_time, avg_turn_around_time)
}
