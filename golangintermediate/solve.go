package golangintermediate

import (
	"fmt"
	"sync"
)

func Sum(d []int, ch chan int, wg *sync.WaitGroup) {
	var result int
	for _, v := range d {
		//hitung
		result += v
		fmt.Println(v)
	}
	ch <- result
	wg.Done()
}
