package main

import (
	"fmt"
)

func main() {

	defer fmt.Println("on defer")
	//os.Exit(0)
	panic("runtime error!")
	fmt.Println("Hello World!")
}
