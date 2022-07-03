package main

import (
	"fmt"
	"sync"
)

type Incrementer struct {
	count int
	sync.Mutex
}

func (s *Incrementer) Increment() {
	s.Lock()
	s.count++
	defer s.Unlock()
}

func main() {
	var wg sync.WaitGroup
	inc := Incrementer{
		count: 0,
	}
	c := 10000000
	wg.Add(c)
	for i := 0; i < c; i++ {
		go func() {
			for i := 0; i < 15; i++ {
				inc.Increment()
			}
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(inc.count)
}
