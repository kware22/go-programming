package main

import (
	"fmt"

	"github.com/Kware22/go-programming/003-packages/cat"
)

func main() {
	fmt.Println("Hello from dog!")

	x := 1
	if x == 1 {
		cat.CatSpeak()
	} else {
		cat.CatFood()
	}

}
