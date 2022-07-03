package main

import (
	"fmt"
	"reflect"
)

func main() {
	ch := make(chan int)
	array := []interface{}{"abc", 10, true, ch}
	fmt.Println("by type switch")
	fmt.Println()
	for _, v := range array {
		getTypeByTypeSwitch(v)
	}
	fmt.Println()
	fmt.Println("by reflection")
	fmt.Println()
	for _, v := range array {
		fmt.Println(getTypeByReflection(v))
	}
}

//type switch
func getTypeByTypeSwitch(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("value is an integer")

	case bool:
		fmt.Println("value is an boolean")
	case chan int:
		fmt.Println("value is a channel")
	case string:
		fmt.Println("value is a string")
	}
}

//с помощью рефлексии
func getTypeByReflection(value interface{}) reflect.Type {
	return reflect.TypeOf(value)
}
