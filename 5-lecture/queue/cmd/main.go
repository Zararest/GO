package main

import (
	"flag"
	"sort"

	"lock-free-queue/pkg/queue"
)

func produce[T any](q *queue.Queue[T], data chan T) {
	for val := range data {
		q.Push(val)
	}
}

func consume[T any](q *queue.Queue[T], out chan T) {
	for {
		val, err := q.Pop()
		if err != nil {
			continue
		}
		out <- val
	}
}

func generateData(n int) []int {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = i
	}
	return arr
}

func checkResults(out chan int, data []int) {
	outData := make([]int, len(data))
	for i := 0; i < len(data); i++ {
		val := <-out
		outData[i] = val
	}

	sort.Ints(outData)
	if !cmp(outData, data) {
		/*for i, data := range outData {
			fmt.Printf("%d: %d\n", i, data)
		}*/
		panic("invalid output array")
	}
}

// FIXME: is there any other way?
func max(lhs, rhs int) int {
	if lhs < rhs {
		return rhs
	}
	return lhs
}

func cmp(lhs, rhs []int) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	// Compare each element
	for i := 0; i < len(lhs); i++ {
		if lhs[i] != rhs[i] {
			return false
		}
	}

	return true
}

func main() {
	var numOfElements int
	var numOfProducers int
	var numOfConsumers int

	flag.IntVar(&numOfElements, "n", 10,
		"Num of elements to generate")
	flag.IntVar(&numOfProducers, "producers", 1,
		"Num of producers")
	flag.IntVar(&numOfConsumers, "consumers", 1,
		"Num of consumers")
	flag.Parse()

	intQueue := queue.Create[int]()
	inputData := generateData(numOfElements)
	sort.Ints(inputData)

	inputChan := make(chan int, numOfElements)

	for _, data := range inputData {
		inputChan <- data
	}

	outChan := make(chan int, numOfElements)
	for i := 0; i < max(numOfProducers, numOfConsumers); i++ {
		if i < numOfProducers {
			go produce(&intQueue, inputChan)
		}
		if i < numOfConsumers {
			go consume(&intQueue, outChan)
		}
	}

	checkResults(outChan, inputData)
}
