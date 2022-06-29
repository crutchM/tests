package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"sync"
	"tests/l0/data"
	"tests/l0/db"
	"tests/l0/repo"
	"tests/l0/server"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	cache := data.NewCache(0, 0)
	dab, err := db.NewConnection()
	if err != nil {
		logrus.Fatal(err)
	}
	cont := db.NewDataBase(dab)
	rep := repo.NewRepository(cont, cache)
	hand := server.NewHandler(rep)
	srv := new(server.Server)
	receiver := server.NewReceiver("test-cluster", rep)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go receiver.Receive()
	srv.Run("8080", hand.InitRoutes())
	wg.Wait()
	fmt.Scanln()

}
