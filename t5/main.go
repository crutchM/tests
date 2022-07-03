package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go sender(ch)
	go receiver(ch)
	time.Sleep(10 * time.Second)
	fmt.Println("Прошло 2 секунды, выключаюсь")
}

func sender(ch chan int) {
	for i := 0; i < 1000; i++ {
		ch <- i
		time.Sleep(time.Millisecond)
	}
	close(ch)
}

func receiver(ch chan int) {
	for {
		select {
		case val := <-ch:
			fmt.Println(val)
		default:
			continue
		}
	}
}
