package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d)
	fmt.Println("wake the fuck up samurai")

}

func main() {
	fmt.Println("start waiting")
	sleep(time.Second * 3)
}
