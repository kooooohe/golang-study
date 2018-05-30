package main

import "fmt"

var plusAlias = plus

func plus(x, y int) int {
	return x + y
}

func main() {

	f := func(x, y int) int { return x + y }

	fmt.Println(f(11, 3))
	fmt.Printf("%T", f)

	//var x func(int, int) int
	println(func(x, y int) int { return x + y }(2, 5))
	plusAlias(1, 2)
}
