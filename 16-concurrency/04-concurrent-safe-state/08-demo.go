/*
modify the below to execute the increment function concurrently
*/
package main

import (
	"fmt"
	"sync"
)

var count int
var mutex sync.Mutex

func main() {
	wg := &sync.WaitGroup{}
	for range 200 {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("Count :", count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	{
		count++
	}
	mutex.Unlock()
}
