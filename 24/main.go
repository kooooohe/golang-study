package main

import (
	"fmt"
)

func main() {
	/**
	println(sum(1, 2, 3))
	println(sum(1, 2, 3, 4, 5))
	s := []int{1, 2, 3}
	println(sum(s...))
	**/
	a := [...]int{1, 2, 3}
	pow(a)
	fmt.Println(a)

	a1 := []int{1, 2, 3}
	pow2(a1)
	fmt.Println(a1)

	var (
		b  [3]int
		b1 []int
	)
	fmt.Println(b)
	fmt.Println(b1 == nil)

	c := [5]int{1, 2, 3, 4, 5}
	s := c[0:2]
	fmt.Println(c)
	fmt.Println(s)
	c[1] = 10
	fmt.Println(s)
	s = append(s, 15, 1, 1, 1)
	c[1] = 20
	fmt.Println(s)
	fmt.Println(c)
}

func pow2(a []int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}
func pow(a [3]int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}

func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}
