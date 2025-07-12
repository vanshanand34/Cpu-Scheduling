package main

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

}
