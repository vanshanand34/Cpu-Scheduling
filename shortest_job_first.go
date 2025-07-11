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

	fmt.Println(processes)

	slices.SortFunc(processes, burstCompare)

	fmt.Println(processes)
	var n int = len(processes)

	var completion_time []float64 = make([]float64, n)

	var waiting_time []float64 = make([]float64, n)
	var turn_around_time []float64 = make([]float64, n)

	waiting_time[0] = 0
	completion_time[0] = processes[0].burst
	turn_around_time[0] = processes[0].burst

	for i := 1; i < n; i++ {
		completion_time[i] = max(completion_time[i-1], processes[i].arrival) + processes[i].burst
		turn_around_time[i] = completion_time[i] - processes[i].arrival
		waiting_time[i] = turn_around_time[i] - processes[i].burst
	}

	var total_waiting_time float64 = 0
	var total_turn_around_time float64 = 0

	for _, wait_time := range waiting_time {
		total_waiting_time += wait_time
	}

	for _, tat := range turn_around_time {
		total_turn_around_time += tat
	}

	avg_waiting_time := total_waiting_time / float64(n)
	avg_turn_around_time := total_turn_around_time / float64(n)

	fmt.Println(waiting_time)
	fmt.Println(turn_around_time)

	fmt.Println("--- Shortest Job First ---")

	fmt.Println("Average Waiting Time: ", avg_waiting_time)
	fmt.Println("Average Turn Around Time: ", avg_turn_around_time)
}
