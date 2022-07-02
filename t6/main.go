package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	data := make(chan int)
	quit := make(chan struct{})
	fmt.Println("Пример остановки с помощью отдельного канала")
	go stopByChannel(quit, data)
	for i := 0; i < 10; i++ {
		data <- i
		if i == 9 {
			quit <- struct{}{}
		}
	}
	fmt.Println("Просто закрытие канала с данными")
	number := stopByChannelBlock()
	fmt.Println(<-number)
	fmt.Println(<-number)
	close(number)
	fmt.Println("Остановка с помощью контекста")
	stopByContext()
}

//с помощью отдельного канала
func stopByChannel(quit chan struct{}, data chan int) {
	for {
		select {
		case <-quit:
			return
		case val := <-data:
			fmt.Println(val)
		}
	}
}

//используем канал для передачи данных и сигнализации об остановке
func stopByChannelBlock() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
		}
	}()
	return ch
}

//с помощью контекста

func stopByContext() {
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				ch <- struct{}{}
				return
			default:
				fmt.Println("loop")
			}

			time.Sleep(time.Second)
		}
	}(ctx)

	go func() {
		time.Sleep(10 * time.Second)
		cancel()
	}()

	<-ch
	fmt.Println("done")

}
