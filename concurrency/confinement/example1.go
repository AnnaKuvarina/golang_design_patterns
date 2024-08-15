package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func process(data int) int {
	time.Sleep(time.Second * 2)
	return data * 2
}

// BEFORE
/*func processData(wg *sync.WaitGroup, result *[]int, data int) {
	defer wg.Done()

	lock.Lock()
	*result = append(*result, process(data))
	lock.Unlock()
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5}
	result := []int{}

	for _, data := range input {
		wg.Add(1)
		go processData(&wg, &result, data)
	}

	wg.Wait()

	fmt.Println(result)
	fmt.Println(time.Since(start))
}*/

func processData(wg *sync.WaitGroup, resultDest *int, data int) {
	defer wg.Done()

	*resultDest = process(data)
}

func main2() {
	start := time.Now()
	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5}
	result := make([]int, len(input))

	for i, data := range input {
		wg.Add(1)
		go processData(&wg, &result[i], data)
	}

	wg.Wait()

	fmt.Println(result)
	fmt.Println(time.Since(start))
}
