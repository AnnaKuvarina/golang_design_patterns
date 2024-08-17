package main

import "fmt"

func sliceToChannel(numbers []int) chan int {
	out := make(chan int)

	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()

	return out
}

func square(in chan int) chan int {
	out := make(chan int)

	go func() {
		for val := range in {
			out <- val * val
		}
		close(out)
	}()

	return out
}

func main() {
	// input
	numbers := []int{1, 2, 3, 4, 5}
	// stage 1
	dataChannel := sliceToChannel(numbers)
	// stage 2
	finalChannel := square(dataChannel)
	// print result
	for n := range finalChannel {
		fmt.Println(n)
	}
}
