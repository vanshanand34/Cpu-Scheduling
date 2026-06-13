package main

import (
	"cmp"
	"slices"
)

func longestJobFirst(processes []Process) {
	// First sort the processes in ascending order of their arrival times
	// Then execute the process with longest burst time (which is available in the queue at current time)

	slices.SortFunc(processes, func(a, b Process) int {
		return cmp.Compare(a.arrival, b.arrival)
	})

	var currentTime = 0.0
	var processExecuted = 0
	var processCount = len(processes)
	var largestBurstIdx = -1
	completed := make([]bool, processCount)

	totalWaitingTime := 0.0
	totalTurnaroundTime := 0.0

	for processExecuted < processCount {

		largestBurstIdx = -1

		// Find process available at current time with largest burst time
		for i := range processCount {
			if completed[i] {
				continue
			}

			if processes[i].arrival > currentTime {
				break
			}

			if largestBurstIdx == -1 || processes[i].burst > processes[largestBurstIdx].burst {
				largestBurstIdx = i
			}
		}

		if largestBurstIdx == -1 {
			currentTime++
			continue
		}

		currProcess := processes[largestBurstIdx]
		completionTime := currentTime + currProcess.burst
		turnAroundTime := completionTime - currProcess.arrival
		waitingTime := turnAroundTime - currProcess.burst

		completed[largestBurstIdx] = true
		processExecuted++
		totalWaitingTime += waitingTime
		totalTurnaroundTime += turnAroundTime
		currentTime = completionTime
	}

	var avgWaitingTime float64 = totalWaitingTime / float64(processCount)
	var avgTurnaroundTime float64 = totalTurnaroundTime / float64(processCount)

	printAlgoResult("Longest Job First Algorithm", avgWaitingTime, avgTurnaroundTime)
}
