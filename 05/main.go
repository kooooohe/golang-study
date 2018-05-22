package main

import (
	_ "fmt"
)

func main() {
	/**
	var a int
	var i int

	var c, d, e string

	var (
		x, y int
		name string
	)
	**/

	var n int
	n = 5
	//n = "sss"
	println(n)

	_ = true
	_ = 4.31
	_ = "aaaa"
	_ = n
	print(one())

	c := one()

	print(c)
}

func one() int {
	return 1
}
