package main

import "fmt"

type Person struct {
	Id   int
	Name string
	Area string
}

func main() {
	p := new(Person)
	fmt.Println(p.Id)
	fmt.Println(p.Name)
	fmt.Println(p.Area)
	i := new(int)
	println(i)
}
