package main

func main() {
	ch := make(chan int, 3)
	ch <- 5
	ch <- 4
	ch <- 3
	/**
	i := <-ch
	i = <-ch
	**/
	i := <-ch
	println(i)
	println(len(ch))
	/**
	if len(ch) > 0 {

	}
	**/
	close(ch)
	ch <- 33
}
