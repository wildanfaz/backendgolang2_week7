package golangintermediate

import (
	"fmt"
	"sync"
)

func FibonacciNumbers(ch chan<- []int, n int, wg *sync.WaitGroup) {
	var result []int
	if n >= 1 {
		result = append(result, 0, 1)
	} else if n >= 0 {
		result = append(result, 0)
	}
	for i := 0; i < n-1; i++ {
		var nextVal = result[i] + result[i+1]
		if nextVal > n {
			break
		}
		result = append(result, nextVal)
	}
	ch <- result
	wg.Done()
}

func OddEven(ch <-chan []int, wg *sync.WaitGroup) {
	num := <-ch
	for _, v := range num {
		if v%2 == 0 {
			fmt.Println(v, "Is Even Number")
		} else {
			fmt.Println(v, "Is Odd Number")
		}
	}
	wg.Done()
}
