package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	dict := make(map[int]int)
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(num int) {
			mutex.Lock()
			dict[num] = rand.Intn(100)
			defer mutex.Unlock()
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(dict)
}
