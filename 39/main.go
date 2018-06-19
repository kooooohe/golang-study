package main

import (
	"fmt"
)

type MyInt int

func (m MyInt) Plus(i int) int {
	return int(m) + i
}

type Point struct{ X, Y int }

func (p *Point) ToString() string {
	return fmt.Sprintf("[%d],[%d]", p.X, p.Y)
}

func main() {
	//println(MyInt(4).Plus(3))
	p := Point{10, 18}
	println(p.ToString())

	f := (*Point).ToString
	// func(*Point) string
	fmt.Printf("%T", f)
	println(f(&Point{X: 7, Y: 11}))
	println((*Point).ToString(&Point{11, 12}))

}
