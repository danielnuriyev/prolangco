package main

import "fmt"
import "time"
import "runtime"

type IntWrapper struct {
	v int
}

func main() {
	
	const size = 1000000
	const rounds = 100
	
	// var numbers [size]int
	var numbers [size]IntWrapper
	
	var min uint64 = 8 * 1024 * 1024 * 1024
	var max uint64 = 0

	memory := make([]uint64, rounds)
	
	var mem runtime.MemStats
	
	start := time.Now()
	
	for round := 0; round < rounds; round++ {
	
		for counter := 0; counter < size; counter++ {
			numbers[counter] = IntWrapper { counter }
		}
		
		runtime.ReadMemStats(&mem)
		
		if mem.Alloc < min {
			min = mem.Alloc
		}
		
		if mem.Alloc > max {
			max = mem.Alloc
		}
		
		memory[round] = mem.Alloc
		
		// fmt.Printf("round: %d\n", mem.Alloc)
		
	}
	
	fmt.Printf("time: %s\n", time.Since(start))
	fmt.Printf("min: %d\n", min)
	var tmp float64 = 0
	for i := 0; i < rounds; i++ {
		tmp += float64(memory[i])
	}
	tmp = tmp / rounds
	fmt.Printf("avg: %f\n", tmp)
	fmt.Printf("max: %d\n", max)
	
}