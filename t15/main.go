package main

import (
	"fmt"
	"unsafe"
)

//var justString string
//func someFunc() {
//  v := createHugeString(1 << 10)
//  justString = v[:100]
//}
//
//func main() {
//  someFunc()
//}

func createHugeString(n int) string {
	return "азм естьм заглушка"
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	//в исходном если просто брать 100 символо из слайса можем словить проблемы с кодировкой, поэтому лучше привести все к слайсу рун
	//также смущает просто использование не входного параметра а глобальной переменной
	justString = string(append([]rune{}, []rune(v)[:100]...))
}

func main() {
	//пример "бед" с utf в случае с b не выведется ожидаемая строка
	a := "фыааврапонелнгл"
	b := a[:4]
	c := string(append([]rune{}, []rune(a)[:4]...))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println(b)
	fmt.Println(c)
}
