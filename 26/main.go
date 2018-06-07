package main

import (
	"fmt"
)

func main() {
	m := make(map[int]string)
	m[1] = "US"
	m[81] = "Japan"
	m[86] = "China"
	fmt.Println(m)
	m1 := map[int]string{1: "Hello",
		5: "kkk",
		6: "aaa",
	}
	fmt.Println(m1)
	m2 := map[int][]int{
		1: []int{1},
		3: []int{1, 2},
		4: []int{1, 2, 3},
	}
	fmt.Println(m2)
	m3 := map[int]map[float64]string{
		1: {3.14: "円周率"},
	}
	fmt.Println(m3)
	m4 := map[int]string{1: "A", 2: "B"}
	fmt.Println(m4[1])
	fmt.Println(m4[10])

	m5 := map[int]int{1: 0}
	s := m5[10]
	fmt.Println(s)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}
	fmt.Println(len(m))
	delete(m, 81)
	fmt.Println(m)
}
