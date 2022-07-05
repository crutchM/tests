package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (s *Human) Speech() {
	fmt.Println("Ку")
}

func (s *Human) GetAge() int {
	return s.Age
}

type Action struct {
	Human
}

func (s *Action) PrintAge() string {
	return fmt.Sprint(s.GetAge())
}

func main() {
	tmp := Action{Human: Human{
		Name: "Alexey",
		Age:  20,
	}}

	tmp.Speech()
	fmt.Println(tmp.PrintAge())
}
