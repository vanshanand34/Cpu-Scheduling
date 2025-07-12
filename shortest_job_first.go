package main

import (
	"cmp"
	"fmt"
	"slices"
)

func shortestJobFirst(processes []Process) {

	burstCompare := func(p1, p2 Process) int {
		return cmp.Compare(p1.burst, p2.burst)
	}

	slices.SortFunc(processes, burstCompare)

	var n int = len(processes)

	var completion_time []float64 = make([]float64, n)
	var waiting_time []float64 = make([]float64, n)
	var turn_around_time []float64 = make([]float64, n)

	waiting_time[0] = 0
	completion_time[0] = processes[0].burst
	turn_around_time[0] = processes[0].burst

	var total_waiting_time float64 = 0
	var total_turn_around_time float64 = 0

	for i := 1; i < n; i++ {
		completion_time[i] = max(completion_time[i-1], processes[i].arrival) + processes[i].burst
		turn_around_time[i] = completion_time[i] - processes[i].arrival
		waiting_time[i] = turn_around_time[i] - processes[i].burst

		total_waiting_time += waiting_time[i]
		total_turn_around_time += turn_around_time[i]
	}

	avg_waiting_time := total_waiting_time / float64(n)
	avg_turn_around_time := total_turn_around_time / float64(n)

	fmt.Println("--- Shortest Job First ---")
	fmt.Println("Average Waiting Time: ", avg_waiting_time)
	fmt.Println("Average Turn Around Time: ", avg_turn_around_time)

}
