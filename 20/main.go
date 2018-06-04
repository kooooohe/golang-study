package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("NumCPU %d\n", runtime.NumCPU())
	fmt.Printf("NumGorutine %d\n", runtime.NumGoroutine())
	fmt.Printf("Version %d\n", runtime.Version())
	go fmt.Println("aaa")
	/**	go sub() //ゴルーチン
	i := 0
	for i < 100 {
		fmt.Println("main Loop")
		i++
	}
	**/

	/**
	go func() {
		k := 0
		for k < 100 {
			println("sub2 loop")
		}
	}()
	**/
}

func sub() {
	i := 0
	for i < 100 {
		fmt.Println("sub Loop")
		i++
	}
}
