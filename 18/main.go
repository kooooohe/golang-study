package main

import (
	"fmt"
)

func main() {

	/**
	LOOP:
		for {
			for {
				continue LOOP
				for {
					break LOOP
				}
			}
		}
		**/
	runDefer()
}

func runDefer() {
	defer func() {
		defer fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()
	fmt.Println("done")
	/**
	file,err:=os.Open("/path/to/tm.txt")
	if err!=nil{
		return
	}
	defer file.Close()
	**/

	/**
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	**/
}
