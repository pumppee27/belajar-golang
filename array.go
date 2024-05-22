package main

import "fmt"

func main() {
	fruits := [...]string{"Apel", "Nangka"}

		for i, j :=range fruits{
			fmt.Println("index", i, "Nama", j)
		}
	
}