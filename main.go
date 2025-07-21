package main

import "fmt"

type Process struct {
	arrival float64
	burst   float64
}

func main() {

	var processes []Process = []Process{
		{0, 1}, {1, 7}, {2, 22}, {3, 3}, {4, 9}, {5, 15}, {6, 26}, {7, 17}, {8, 8}, {9, 19},
	}

	firstComeFirstServe(processes)
	shortestJobFirst(processes)
	longestJobFirst(processes)
	highestResponseRatioNext(processes)

}

func printAlgoResult(algo_name string, avg_wait_time, avg_tat_time float64) {
	fmt.Println()
	fmt.Println("--- ", algo_name, " ---")
	fmt.Println("Average waiting time: ", avg_wait_time)
	fmt.Println("Average turn around time: ", avg_tat_time)
	fmt.Println()
}
