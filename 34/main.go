package main

import (
	"fmt"
)

func main() {
	//type T0 int
	//type T1 int

	/**
	t0 := T0(5)
	i0 := int(t0)

	t1 := T1(8)
	i1 := int(t1)

	t0 = t1
	**/
	type Point struct {
		X, Y int
	}

	var pt Point
	pt.X = 10
	fmt.Println(pt)
	pt2 := Point{1, 2}
	fmt.Println(pt2)
	pt3 := Point{Y: 2, X: 33}
	fmt.Println(pt3)
	type Person struct {
		ID   uint
		name string
		部署   string
		_    int
	}
	p := Person{ID: 17, 部署: "営業", name: "aa"}
	fmt.Println(p)

	type Feed struct {
		Name   string
		Amount uint
	}

	type Animal struct {
		Name string
		Feed
	}

	a := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 10,
		},
	}
	fmt.Println(a)
	println(a.Name)
	println(a.Feed.Name)
	println(a.Amount)

	type T0 struct {
		Name1 string
	}
	type T1 struct {
		T0
		Name2 string
	}
	type T2 struct {
		T1
		Name3 string
	}
	t := T2{T1: T1{T0: T0{Name1: "X"}, Name2: "Y"}, Name3: "Z"}
	println(t.Name1)
	println(t.Name2)
	println(t.Name3)

	type T22 struct {
		int
		float64
		string
	}
	t22 := T22{1, 3.14, "文字列"}
	println(t22.int)
	println(t22.string)
	println(t22.float64)
}
