package main

func main() {
	ch := make(chan int, 3)
	ch <- 5
	ch <- 4
	ch <- 3
	i := <-ch
	i = <-ch
	close(ch)
	i, ok := <-ch
	println(i)
	println(ok)
	i, ok = <-ch
	println(i)
	println(ok)
	//println(len(ch))
	/**
	if len(ch) > 0 {

	}
	**/
	//ch <- 33
}
