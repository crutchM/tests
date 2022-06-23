package main

import (
	"github.com/sirupsen/logrus"
	"tests/l0"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	cache := l0.NewCache(0, 0)
	db, err := l0.NewConnection()
	if err != nil {
		logrus.Fatal(err)
	}
	cont := l0.NewDataBase(db)

}
