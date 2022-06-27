package publisher

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

type Pub struct {
	con stan.Conn
}

func NewStan() *Pub {
	con, err := stan.Connect("test-cluster", "publisher", stan.NatsURL("nats://192.168.0.104:4422"))
	if err != nil {
		logrus.Fatal(err.Error())
	}
	return &Pub{con: con}
}

func (s *Pub) Publish(value interface{}) error {
	err := s.con.Publish("receiver", []byte(fmt.Sprint(value)))
	if err != nil {
		return err
	}

	return nil
}
