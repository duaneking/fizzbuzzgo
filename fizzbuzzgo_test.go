package main

import (
	"fmt"
	"testing"
)

func assert(t *testing.T, what string, with string) {
	if what != with {
		t.Errorf("%s != %s", what, with)
	}
}

func TestFizzBuzzBasic(t *testing.T) {
	in := make(chan int, 3)
	out := make(chan string, 3)

	go ChannelFizzBuzz(in, out)

	in <- 15
	in <- 3
	in <- 5

	expectFB := <-out
	assert(t, expectFB, Fizz+Buzz)

	expectF := <-out
	assert(t, expectF, Fizz)

	expectB := <-out
	assert(t, expectB, Buzz)
}

func TestFizzBuzzChannelTableDriven(t *testing.T) {
	var tests = []struct {
		index  int
		expect string
	}{
		{15, Fizz + Buzz},
		{3, Fizz},
		{5, Buzz},
	}

	for _, tt := range tests {
		in := make(chan int, 4)
		out := make(chan string, 4)

		testname := fmt.Sprintf("%d,%s", tt.index, tt.expect)
		t.Run(testname, func(t *testing.T) {
			go ChannelFizzBuzz(in, out)

			in <- tt.index

			expected := <-out

			assert(t, expected, tt.expect)
		})
	}
}

func TestFizzBuzzOperationTableDriven(t *testing.T) {
	var tests = []struct {
		index  int
		expect string
	}{
		{15, Fizz + Buzz},
		{3, Fizz},
		{5, Buzz},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%s", tt.index, tt.expect)
		t.Run(testname, func(t *testing.T) {
			expected := FizzBuzzOperation(tt.index)
			assert(t, expected, tt.expect)
		})
	}
}

// BenchmarkFizzBuzz is important because using channels actually slows this down a lot.
func BenchmarkFizzBuzz(b *testing.B) {
	order := make(chan int)
	message := make(chan string)

	go ChannelFizzBuzz(order, message)

	for index := 1; index < 101; index++ {
		order <- index
		s := <-message
		fmt.Println(s)
	}
}
