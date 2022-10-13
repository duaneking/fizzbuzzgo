package main

import (
	"fmt"
	"runtime"
)

/* Constants */
const Fizz string = "Fizz"
const Buzz string = "Buzz"

// FizzBuzz is the simple, fast FizzBuzz that your momma probably warned you about.
func FizzBuzz(max int) {
	for i := 1; i < max; i++ {
		fbo := FizzBuzzOperation(i)

		fmt.Print(fbo)
	}
}

// FizzBuzzOperation is the actual logic for FizzBuzz, extracted as a unit testable function.
func FizzBuzzOperation(i int) string {
	switch {
	case i%15 == 0:
		return Fizz + Buzz
	case i%3 == 0:
		return Fizz
	case i%5 == 0:
		return Buzz
	default:
		return fmt.Sprintf("%d", i)
	}
}

// ChannelFizzBuzz is the version of FizzBuzz that happens when you
// use Go and want to show how channels work but want the original
// FizzBuzz to be well factored and easily testable.
func ChannelFizzBuzz(input chan int, output chan string) {
	for {
		index := <-input

		output <- FizzBuzzOperation(index)
	}
}

func main() {
	// Scale to system.
	cpus := runtime.NumCPU()

	// Use buffered channels, leveraging system cores.
	input := make(chan int, cpus)
	results := make(chan string, cpus)

	// Set up channels
	go ChannelFizzBuzz(input, results)

	// Let it work. Serialize output for sanity.
	for index := 1; index < 101; index++ {
		input <- index
		value := <-results
		fmt.Println(value)
	}
}
