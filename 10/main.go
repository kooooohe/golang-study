package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	//interface{}
	/**
	var x interface{}
	x = "ss"
	fmt.Println(x)
	x = 1
	fmt.Println(x)
	x = '山'
	fmt.Println(x)
	x = [...]uint8{1, 2, 3}
	fmt.Println(x)
	**/

	//var y, z interface{}
	//y, z = 1, 2
	//z := y + z
	a := puls(1, 5)
	println(a)
	//var a,b  int
	hello()
	q, _ := div(19, 7)
	fmt.Printf("%v,", q)

	/**
	_, err = div(19, 7)
	if err!={
		//エラー処理
	}
	**/
	//fmt.Printf("%v,", q)
	println(doSomething())
}

func do3(_, _ int) int {
	return 1
}
func doSomething2() (a, b int) {
	/**
	var a,b int
	return a, b
	**/
	return
}
func doSomething() (a int) {
	/**
	var a int
	return a
		**/
	return
}

func puls(x, y int) int {
	return x + y
}

func hello() {
	println("hello")
	return
}

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}
