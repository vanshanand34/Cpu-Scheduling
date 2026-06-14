// Priority Scheduling algorithm implementation
// process with highest priority gets executed first
// processes with same priority gets exectuted in the order of their arrival
// on a First Come First Server basis

// Lower number = Higher Priority
package main

import (
	"cmp"
	"slices"
)

func priorityScheduling(processes []Process) {
	var n int = len(processes)

	slices.SortFunc(processes, func(a, b Process) int {
		return cmp.Compare(a.arrival, b.arrival)
	})

	currentTime := 0.0
	processCompleted := 0
	executed := make([]bool, n)
	bestPriorityIdx := -1
	totalWaitingTime := 0.0
	totalTurnaroundTime := 0.0

	for processCompleted < n {
		// Find the process present in current time with highest priority (least priority number)
		bestPriorityIdx = -1

		for i := range n {

			if executed[i] {
				continue
			}

			currProcess := processes[i]
			if currProcess.arrival > currentTime {
				break
			}

			if bestPriorityIdx == -1 || currProcess.priority < processes[bestPriorityIdx].priority {
				bestPriorityIdx = i
			}

		}

		if bestPriorityIdx == -1 {
			// Find process with least arrival time that is not executed yet, in case no process has not arrived yet at currentTime
			for i := range n {
				if !executed[i] {
					currentTime = processes[i].arrival
					break
				}
			}
			continue
		}

		// Execute process with index bestPriorityIdx
		processToExecute := processes[bestPriorityIdx]
		completionTime := currentTime + processToExecute.burst
		turnaroundTime := completionTime - processToExecute.arrival
		waitingTime := turnaroundTime - processToExecute.burst

		totalWaitingTime += waitingTime
		totalTurnaroundTime += turnaroundTime

		currentTime = completionTime
		executed[bestPriorityIdx] = true
		processCompleted++

	}

	avgWaitingTime := totalWaitingTime / float64(n)
	avgTurnaroundTime := totalTurnaroundTime / float64(n)

	printAlgoResult("Non Preemptive Priority Scheduling Algorithm", avgWaitingTime, avgTurnaroundTime)
}
