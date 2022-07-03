package main

import (
	"fmt"
	"math/rand"
	"time"
)

type client struct {
	email string
	payment
}

type bankAdapter struct {
}

func (s *bankAdapter) Send(fromEmail, toEmail string, amount int) {
	from := rand.Int()
	to := rand.Int()
	fmt.Printf("Send %d from %d to %d at %v via bank", amount, from, to, time.Now().String())
}

type payPalAdapter struct {
}

func (s *payPalAdapter) Send(fromEmail, toEmail string, amount int) {
	from := "+79877902767"
	to := "+79022139176"
	fmt.Printf("Send %d  from %s to %s via paypal", amount, from, to)
}

type payment interface {
	Send(fromEmail, toEmail string, amount int)
}

func main() {
	c1 := client{
		email:   "exmpl@gmail.com",
		payment: &bankAdapter{},
	}

	c2 := client{
		email:   "test@gmail.com",
		payment: &payPalAdapter{},
	}
	c1.Send(c1.email, c2.email, 10)
	fmt.Println()
	c2.Send(c2.email, c1.email, 10)
}
