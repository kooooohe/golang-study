package main

import (
	"fmt"
)

func main() {
	type Base struct {
		Id    int
		Owner string
	}
	type A struct {
		Base
		Name string
		Area string
	}
	type B struct {
		Base
		Title  string
		Bodies []string
	}
	a := A{
		Base: Base{17, "Taro"},
		Name: "Taro",
		Area: "Tokyo",
	}
	b := B{
		Base:   Base{71, "jiro"},
		Title:  "no title",
		Bodies: []string{"A", "B"},
	}
	println(a.Id)
	println(a.Owner)
	println(b.Id)
	println(b.Owner)

	/**
	struct{
		T1
		*T2    //T2
		P.T3   //T3
		*P.T4  //T4
	}
	**/
	/**
	type T3 struct {
		T1
	}
	type T1 struct {
		T3
	}
	**/

	s := struct{ X, Y int }{X: 1, Y: 23}
	showStruct(s)
	type Point struct {
		X, Y int
	}
	p := Point2{X: 3, Y: 10}
	showStruct(p)
}

type Point2 struct {
	X, Y int
}

func showStruct(s struct{ X, Y int }) {
	fmt.Println(s)
}
