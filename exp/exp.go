package main

import (
	"fmt"
	"rl/glue"
)

// const max_steps = 10000

var (
	num_episodes           = 8000
	max_steps              = 10000
	done         chan bool = make(chan bool)
)

func single_run(run int) {
	fmt.Printf("run number %d\n\n", (run + 1))
	rl.Init()
	// fmt.Println()

	for episode := 0; episode < num_episodes; episode++ {
		rl.Episode(max_steps)
	}
	rl.Cleanup()
	done <- true
}

func main() {

	num_runs := 10
	runs := make([]float64, num_runs)
	runs[0] = 0.1
	runs[1] = 0.2
	runs[2] = 0.3
	for _, value := range runs {
		fmt.Println(value)
	}

	for run := 0; run < num_runs; run++ {
		go single_run(run)
	}
	for run := 0; run < num_runs; run++ {
		<-done
	}
}
