package main

import (
	"fmt"
	"time"
)

type ControlMessage int

const (
	DoExit = iota
	ExitOk
)

type Job struct {
	data int
}

type Result struct {
	result int
	job    Job
}

func doubler(jobs <-chan Job, results chan<- Result, control chan ControlMessage) {
	for {
		select {
		case msg := <-control: // read from control
			switch msg {
			case DoExit:
				fmt.Println("Exit goroutine")
				control <- ExitOk // write to control
				return
			default:
				panic("unhandled control message")
			}
		case job := <-jobs: // read from jobs
			results <- Result{result: job.data * 2, job: job} // write to results
		}
	}
}

func main() {
	jobs := make(chan Job, 50)
	results := make(chan Result, 50)
	control := make(chan ControlMessage)

	go doubler(jobs, results, control)

	for i := 0; i < 30; i++ {
		jobs <- Job{i}
	}

	for {
		select {
		case result := <-results:
			fmt.Println(result)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("timed out")
			control <- DoExit

			// wait for goroutine to respond
			// the reason for that is we don't want to exit before goroutine finishes
			<-control

			fmt.Println("program exit")
			return
		}
	}
}
