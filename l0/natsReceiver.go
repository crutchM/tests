package l0

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"time"
)

type Receiver struct {
	con *nats.Conn
}

func NewReceiver(token string) *Receiver {
	nc, err := nats.Connect(token)
	if err != nil {
		logrus.Error(err)
		time.Sleep(5 * time.Second)
		NewReceiver(token)
	}
	return &Receiver{con: nc}
}

func (s *Receiver) Receive() {

}
