package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func main() {
	stopChan := make(chan os.Signal, 1)
	dataChan := make(chan interface{})
	signal.Notify(stopChan, os.Interrupt)
	workers := 0
	fmt.Scanln(&workers)

	for i := 0; i < workers; i++ {
		go worker(i, dataChan, stopChan)
	}

	for i := 0; i < 3000; i++ {
		dataChan <- randString(10)
	}
	time.Sleep(60 * time.Second)
}

func randString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func worker(index int, data chan interface{}, stop chan os.Signal) {
	for {
		select {
		case value := <-data:
			fmt.Println("Worker №", index, " Data: ", value)
		case <-stop:
			os.Exit(1)
		default:
			time.Sleep(2 * time.Second)
		}
	}
}
