package main

import (
	"cmp"
	"fmt"
	"slices"
)

func longestJobFirst(processes []Process) {
	// sort the processes on the basis of burst time of processes
	// in decreasing order
	burstComp := func(p1, p2 Process) int {
		return cmp.Compare(p2.burst, p1.burst)
	}

	n := len(processes)

	slices.SortFunc(processes, burstComp)

	fmt.Println(processes)

	var waiting_time []float64 = make([]float64, n)
	var completion_time []float64 = make([]float64, n)
	var turn_around_time []float64 = make([]float64, n)

	waiting_time[0] = 0
	completion_time[0] = processes[0].burst
	turn_around_time[0] = processes[0].burst

	var total_tat_time float64 = 0
	var total_wait_time float64 = 0

	for i := 1; i < n; i++ {

		start_time := max(completion_time[i-1], processes[i].arrival)
		completion_time[i] = start_time + processes[i].burst
		turn_around_time[i] = completion_time[i] - processes[i].arrival
		waiting_time[i] = turn_around_time[i] - processes[i].burst

		total_wait_time += waiting_time[i]
		total_tat_time += turn_around_time[i]
	}

	avg_wait_time := total_wait_time / float64(n)
	avg_tat_time := total_tat_time / float64(n)

	fmt.Println("--- Longest Job First Algorithm ---")
	fmt.Println("Average waiting time: ", avg_wait_time)
	fmt.Println("Average turn around time: ", avg_tat_time)

}
