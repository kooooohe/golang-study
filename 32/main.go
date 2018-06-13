package main

import (
	"fmt"
)

func main() {
	p := &[3]int{1, 2, 3}
	//println((*p)[1])
	println(p[1])
	fmt.Println(cap(p))
	fmt.Println(len(p))

	for _, v := range p {
		println(v)
	}
}
