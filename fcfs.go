package main

import "fmt"

func fcfs(arrival_time []float64, burst_time []float64) {
	// find waiting times and turn around time
	var n int = len(arrival_time)

	// completion time: Time at which process execution completed
	var completion_time []float64 = make([]float64, n)

	// waiting time = service time - arrival time
	// time for which a process has to wait for CPU to start its execution
	var waiting_time []float64 = make([]float64, n)
	var turn_around_time []float64 = make([]float64, n)

	waiting_time[0] = 0
	completion_time[0] = burst_time[0]
	turn_around_time[0] = burst_time[0]
	waiting_time[0] = 0

	for i := 1; i < n; i++ {
		// fmt.Print(i, " ")
		completion_time[i] = max(completion_time[i-1], arrival_time[i]) + burst_time[i]
		turn_around_time[i] = completion_time[i] - arrival_time[i]
		waiting_time[i] = turn_around_time[i] - burst_time[i]
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

	fmt.Println("Average Waiting Time: ", avg_waiting_time)
	fmt.Println("Average Turn Around Time: ", avg_turn_around_time)
}

func main() {

	// arrival time: Time at which process is placed in ready queue (are ready to be executed)
	// burst time: Time required by a process to execute completely

	// arrival times and burst times of processes

	var arrival_time []float64 = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var burst_time []float64 = []float64{1, 7, 22, 3, 9, 15, 26, 17, 8, 19}

	fcfs(arrival_time, burst_time)

	fmt.Println("Hello, World!")
}
