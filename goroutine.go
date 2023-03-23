package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
var m sync.Mutex

func main() {
	var bisa interface{}="[bisa1 bisa2 bisa3]" 
	var coba interface{}="[coba1 coba2 coba3]" 

	for i := 1; i <= 4; i++ {
	wg.Add(2)
	m.Lock()
	go goroutine1(bisa,i)
	m.Lock()
	go goroutine2(coba,i)
	}
	wg.Wait()
}

func goroutine1(data interface{}, index int) {
	fmt.Println(data,index)
	m.Unlock()
	wg.Done()
}
func goroutine2(data interface{}, index int) {
	
	fmt.Println(data,index)	
	m.Unlock()
	wg.Done()
}