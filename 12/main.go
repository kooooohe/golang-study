package main

import (
	"fmt"
)

func main() {
	//f := returnFunc()
	//f()
	//returnFunc()()

	callFuntion(func(x int) {
		fmt.Println("hello call function")
	})
}

func callFuntion(f func(x int)) {
	/**
	非同期処理
		**/
	f(1)
}

func returnFunc() func() {
	return func() {
		fmt.Println("Im a function")
	}
}
