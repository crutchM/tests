package main

import (
	"fmt"
	"time"
)

func main() {
	go sender()
	time.Sleep(2 * time.Second)
	fmt.Println("Прошло 2 секунды, выключаюсь")
}

func sender() {
	ch := make(chan int)
	defer close(ch)
	go receiver(ch)
	for i := 0; i < 1000; i++ {
		ch <- i
		time.Sleep(time.Second / 10)
	}
}

func receiver(ch chan int) {
	for {
		select {
		case val := <-ch:
			fmt.Println(val)
		}
	}
}
