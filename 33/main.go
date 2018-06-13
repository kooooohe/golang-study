package main

type Callback func(i int) int

func Sum(ints []int, c Callback) int {
	var sum int
	for _, i := range ints {
		sum += i
	}
	return c(sum)
}

func main() {
	c := Sum([]int{1, 2, 3}, func(a int) (b int) {
		return
	})
	println(c)
	type MyInt int

	var n1 MyInt = 5
	n2 := MyInt(7)
	println(n1)
	println(n2)

	type (
		IntPair     [2]int
		Strings     []string
		areaMap     map[string][2]float64
		aaa         [3][3][3]int
		IntsChannel chan []int
	)
	_ = IntPair{1, 2}
	_ = Strings{"Apple", "Banana"}
	_ = areaMap{"Tokyo": {35.1111, 12.111}}
	_ = make(IntsChannel)

}
