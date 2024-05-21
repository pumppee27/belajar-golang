package main

import (
	"fmt"
	"reflect"
)

func main() {
	bilangan_bulat := 10
	float := 9.7
	nama := "Kusuma"
	boolean := true

	fmt.Println(reflect.TypeOf(bilangan_bulat))
	fmt.Println(reflect.TypeOf(float))
	fmt.Println(reflect.TypeOf(nama))
	fmt.Println(reflect.TypeOf(boolean))
}