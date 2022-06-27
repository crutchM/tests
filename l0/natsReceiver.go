package l0

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"time"
)

type Receiver struct {
	con  stan.Conn
	repo *Repository
}

func NewReceiver(token string, repo *Repository) *Receiver {
	nc, err := stan.Connect("test-cluster", "subscriber", stan.NatsURL("nats://192.168.0.104:4422"))
	if err != nil {
		logrus.Error(err)
		time.Sleep(5 * time.Second)
		NewReceiver(token, repo)
	}
	return &Receiver{con: nc, repo: repo}
}

func (s *Receiver) Receive() {
	var item Order
	sub, err := s.con.Subscribe("l0", func(msg *stan.Msg) {
		err := json.Unmarshal(msg.Data, &item)
		if err != nil {
			logrus.Info(err.Error())
			return
		}
		s.repo.Write(item)
	},
		stan.DeliverAllAvailable())
	if err != nil {
		sub.Close()
		s.con.Close()
	}
}
