package main

import (
	"fmt"
	"sync"

	"github.com/wildanfaz/backendgolang2_week7/golangintermediate"
)

func main() {
	wg := &sync.WaitGroup{}
	mt := &sync.RWMutex{}
	fmt.Println("Task Solve")
	a := []int{7, 10, 2, 34, 33, -12, -8, 4}
	chn := make(chan int)

	wg.Add(2)
	go golangintermediate.Sum(a[:len(a)/2], chn, wg)
	fmt.Println("Total =", <-chn)
	go golangintermediate.Sum(a[len(a)/2:], chn, wg)
	fmt.Println("Total =", <-chn)
	wg.Wait()

	wg.Add(4)
	fmt.Println("\nTask 1")
	result := golangintermediate.Limiter{N: 40}
	go result.Prime(wg)
	go result.Odd(wg)
	go result.Even(wg)
	go result.Fibonacci(wg)
	wg.Wait()

	wg.Add(2)
	fmt.Println("\nTask 2")
	fibo := make(chan []int)
	go golangintermediate.FibonacciNumbers(fibo, 40, wg)
	go golangintermediate.OddEven(fibo, wg)
	wg.Wait()

	fmt.Println("\nTask 3")
	food := make(chan []string, 2)
	firstMenu := []string{"Soto Ayam", "Es Teh Manis"}
	secondMenu := []string{"Soto Babat", "Es Jeruk Manis"}
	var allMenu [][]string
	allMenu = append(allMenu, firstMenu, secondMenu)
	for i := 0; i < len(allMenu); i++ {
		wg.Add(2)
		mt.Lock()
		go golangintermediate.PayFood(food, allMenu, wg, mt)
		mt.RLock()
		go golangintermediate.DetailPayment(food, wg, mt)
	}
	wg.Wait()
}
