package main

import (
	"fmt"
)

func main() {
	//チャネル
	/**
	var (
		ch  chan int
		ch1 <-chan int
		ch2 chan<- int
	)
	**/
	ch := make(chan int, 3)
	go receiver(ch)
	i := 0
	for i < 10000 {
		ch <- i
		i++
	}
}

func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}
