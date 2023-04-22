package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p Person) IsAdult() bool {

	if p.age > 18 {
		fmt.Println(p.age)
	}
	return true
}

func main() {
	p := Person{"vamsi", 30}
	p.IsAdult()

}
