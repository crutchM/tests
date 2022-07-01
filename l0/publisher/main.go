package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("test-cluster", "stan-pub")
	if err != nil {
		log.Fatal("can not connect to stan")
	}
	defer sc.Close()
	file, _ := ioutil.ReadFile("model.json")

	_ = sc.Publish("l0-task", file)
	fmt.Println("успешно отправлено")
	time.Sleep(5 * time.Second)
}
