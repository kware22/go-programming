package main

//Methods
import (
	"fmt"
)

type Person struct {
	first string
	last  string
	age   int
}

type human interface {
	speak()
}

func (p *Person) speak() {
	fmt.Println("Hello ", p.first)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := Person{"Kendra", "Ware", 28}
	saySomething(&p)
	p.speak()

}
