package main

import (
	"fmt"
)

func main() {
	//	var b bool
	//	b = true

	// b := false

	/**

	int8
	int16
	int32
	int64

	uint8
	16
	32
	64

	byte
	float32
	64


	**/

	//n := uint(17)
	//i:=1.0
	//i1 := float64(i)

	zero := 0.0
	pinf := 1.0 / zero
	ninf := -1.0 / zero
	nan := zero / zero

	fmt.Printf("%v %v %v \n", pinf, ninf, nan)

	c := '松'
	fmt.Printf("%v", c)
}
