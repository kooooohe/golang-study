package main

import (
	"fmt"
)

func main() {
	/**
	i := 0
	for {
		println(i)
		if i == 100 {

			break
		}
	}
	for i < 100 {
		println(i)
	}
	k := 0
	for k < 100 {
		k++
		continue
	}
	**/
	fruits := [3]string{"apple", "banaan", "Cherry"}
	for i, s := range fruits {
		println(i)
		println(s)
	}
	for i, r := range "ABC" {
		fmt.Printf("%d %d \n", i, r)
	}

	switch n := 3; n {
	case 1, 2:
		println("1 or 2")
	case 3, 4:
		println("3 or 4")
		fallthrough
	default:
		println("unknown")
	}

	var x interface{} = 3
	//i, isInt := x.(float32)
	//println(i)
	//println(isInt)
	switch v := x.(type) {
	case bool:
		print("bool")
		print(v)
	case int, uint:
		print("int, uint")
		//print(v * v)
	case string:
		print("default")
		print(v)
	default:
		print("default")
	}

	println("A")
	goto L
	println("B")
L:
	print("C")
	for {
		for {
			for {
				goto DONE
			}
		}
	}
DONE:
}
