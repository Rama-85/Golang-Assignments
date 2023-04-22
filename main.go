package main

import "fmt"

type Book struct {
	title  string
	author string
	price  int
}

func (b *Book) DiscountPrice(discount int) int {
	dp := b.price * discount / 100
	ap := b.price - dp
	return ap
}

func main() {
	b := Book{"Introduction to Golang", "Krishna", 500}
	//b.DiscountPrice(10)
	fmt.Println(b.DiscountPrice(10))
}
