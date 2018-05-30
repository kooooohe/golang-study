package main

func main() {
	//f, a := later()
	f := later()
	println(f("Golang"))
	println(f("is"))
	println(f("awesome!"))
}

func later() (x func(string) string, s1 string) {
	var store string
	x = func(next string) string {
		s := store
		store = next
		return s
	}
	s1 = "Aa"
	return
}
