package main

import "fmt"

func getCompeleteName() (firstName, lastName string) {
	firstName = "Cahya"
	lastName = "Kusuma"

	return firstName, lastName

}

func main() {
	namaDepan, namaBelakang := getCompeleteName()

	fmt.Println(namaDepan, namaBelakang)
}