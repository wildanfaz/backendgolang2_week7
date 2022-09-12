package golangintermediate

import (
	"fmt"
	"strings"
	"sync"
)

var balance int = 100000
var price int
var total int
var menu int = 1
var i int = 0

func PayFood(food chan []string, allMenu [][]string, wg *sync.WaitGroup, mt *sync.RWMutex) {
	if balance-total > 0 {
		fmt.Printf("Balance = Rp.%d\n", balance-total)
	}
	var list []string
	food <- allMenu[0]
	food <- allMenu[1]
	list = append(list, allMenu[i][0], allMenu[i][1])
	i++
	for _, v := range list[len(list)-2:] {
		if strings.ToLower(v) == "soto ayam" {
			total += 12000
		} else if strings.ToLower(v) == "soto babat" {
			total += 15000
		} else if strings.ToLower(v) == "es teh manis" {
			total += 4000
		} else if strings.ToLower(v) == "es jeruk manis" {
			total += 5000
		}
	}
	if menu == 2 {
		close(food)
	}
	mt.Unlock()
	wg.Done()
}

func DetailPayment(food chan []string, wg *sync.WaitGroup, mt *sync.RWMutex) {
	menu1, menu2 := <-food, <-food
	if menu == 1 {
		price += total
		fmt.Println("Menu 1 =", menu1)
		fmt.Printf("Price = Rp.%d\n", price)
	} else {
		fmt.Println("Menu 2 =", menu2)
		fmt.Printf("Price = Rp.%d\n", total-price)
		fmt.Printf("Total Cost = Rp.%d\n", total)
	}
	if balance-total <= 0 {
		fmt.Printf("Insufficient Balance\n\n")
	} else {
		fmt.Printf("Your Remaining Balance = Rp.%d\n\n", balance-total)
	}
	menu++
	mt.RUnlock()
	wg.Done()
}
