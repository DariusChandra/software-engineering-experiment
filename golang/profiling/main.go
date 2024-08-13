package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func busyWork() {
	sum := 0
	for i := 0; i < 1e7; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)
}

func main() {
	// Start CPU profiling
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("could not create CPU profile:", err)
		return
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("could not start CPU profile:", err)
		return
	}
	defer pprof.StopCPUProfile()

	// Simulate some work
	busyWork()

	// Sleep to allow some time for the profiling
	time.Sleep(2 * time.Second)
}
