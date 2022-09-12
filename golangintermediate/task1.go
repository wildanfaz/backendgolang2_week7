package golangintermediate

import (
	"fmt"
	"sync"
)

type Limiter struct {
	N int
}

func (num *Limiter) Prime(wg *sync.WaitGroup) {
	var result []int
	if num.N >= 5 {
		result = append(result, 2, 3, 5)
	} else if num.N >= 3 {
		result = append(result, 2, 3)
	} else if num.N >= 2 {
		result = append(result, 2)
	}
	for i := 0; i <= num.N; i++ {
		if i <= 1 || i%2 == 0 || i%3 == 0 || i%5 == 0 {
			continue
		} else {
			result = append(result, i)
		}
	}
	fmt.Println("Prime =", result)
	wg.Done()
}

func (num *Limiter) Odd(wg *sync.WaitGroup) {
	var result []int
	for i := 0; i <= num.N; i++ {
		if i%2 == 1 {
			result = append(result, i)
		}
	}
	fmt.Println("Odd =", result)
	wg.Done()
}

func (num *Limiter) Even(wg *sync.WaitGroup) {
	var result []int
	for i := 0; i <= num.N; i++ {
		if i%2 == 0 && i > 0 {
			result = append(result, i)
		}
	}
	fmt.Println("Even =", result)
	wg.Done()
}

func (num *Limiter) Fibonacci(wg *sync.WaitGroup) {
	var result []int
	if num.N >= 1 {
		result = append(result, 0, 1)
	} else if num.N >= 0 {
		result = append(result, 0)
	}
	for i := 0; i < num.N-1; i++ {
		var nextVal = result[i] + result[i+1]
		if nextVal > num.N {
			break
		}
		result = append(result, nextVal)
	}
	fmt.Println("Fibonacci =", result)
	wg.Done()
}
