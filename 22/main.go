package main

import (
	"fmt"
)

func main() {
	/**
	slice,map,channel

	make()
	**/

	//スライス
	s := make([]int, 5, 10)
	fmt.Println(s)
	fmt.Println(len(s))

	a := [...]int{1, 2, 3}

	fmt.Println(len(a))
	s1 := []int{1, 2, 3, 4}
	fmt.Println(s1)

	a1 := [5]int{1, 2, 3, 4, 5}
	s2 := a1[0:2]
	fmt.Println(s2)

	fmt.Println(a1[len(a1)-2])

	s3 := "ABCDE"[1:3]
	fmt.Println(s3)

	s4 := "あいうえお"[3:9]
	fmt.Println(s4)

	s5 := []int{1, 2, 3}
	s5 = append(s5, 4)
	fmt.Println(s5)

	s5 = append(s5, 5, 6, 7)
	fmt.Println(s5)

	s6 := []int{1, 2, 3}
	s7 := []int{2, 3, 4}
	s8 := append(s6, s7...)
	fmt.Println(s8)
}
