package server

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"tests/l0/data"
	"tests/l0/repo"
	"time"
)

type Receiver struct {
	con  stan.Conn
	repo *repo.Repository
}

func NewReceiver(token string, repo *repo.Repository) *Receiver {
	nc, err := stan.Connect("test-cluster", "subscriber", stan.NatsURL("nats://192.168.0.104:4422"))
	if err != nil {
		logrus.Error(err)
		time.Sleep(5 * time.Second)
		NewReceiver(token, repo)
	}
	return &Receiver{con: nc, repo: repo}
}

func (s *Receiver) Receive() {

	var item data.Order
	sub, err := s.con.Subscribe("l0-task", func(msg *stan.Msg) {
		err := json.Unmarshal(msg.Data, &item)
		if err != nil {
			logrus.Info(err.Error())
			return
		}
		s.repo.Write(item)
		logrus.Info("объект успешно получен из канала и записан в базу")
	},
		stan.StartWithLastReceived())
	if err != nil {
		sub.Close()
		s.con.Close()
	}
	fmt.Scanln()
}
