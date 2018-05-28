package main

import (
	"fmt"
)

func main() {
	/**
		s := "hello \n world"
		fmt.Println(s)
		s1 := `
		Goの
		文字列リテラル
		による\n\n
		複数行の
		文字列
		`
		fmt.Println(s1)
	**/
	var a1 int
	fmt.Println(a1)
	var a2 int
	fmt.Println(a2)

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v", a)
	c := [5]int{}
	fmt.Printf("%v", c)
	c1 := [5]int{1, 2, 3}
	fmt.Printf("%v", c1)

	var c2 [5]int
	fmt.Println(c2)

	c3 := [3]string{}
	fmt.Println(c3)

	c4 := [3]bool{}
	fmt.Println(c4)

	c5 := [0]int{}
	fmt.Println(c5)

	d := [...]int{1, 2, 3}
	fmt.Println(d)

	d[0] = 0
	d[2] = 0
	fmt.Println(d)

	d1 := [...]uint8{1, 2, 3}
	d1[0] = 255

	var (
		d2 [5]int
		//d3 [5]uint
		d3 [5]int
	)
	d2 = d3
	fmt.Println(d2)

	e := [3]int{1, 2, 3}
	e1 := [3]int{1, 5, 6}
	e = e1

	fmt.Println(e)
	e[0] = 0
	e[1] = 0
	fmt.Println(e)
	fmt.Println(e1)

	//interface{}
	var x interface{}
	fmt.Println(x)
}
