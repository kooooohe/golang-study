package main

// 大文字にすると外部から呼び出せる(定数とは関係なし)
const X = 1

func main() {
	//fmt.Println()

	const (
		x = 1
		Z
		Y
		XX, YY = 33, 33
	)
	const (
		x1 = 1
		Z1
		Y1
		Xx1 = "あ"
		Xx2
	)
	const (
		X2 = 2
		Y2 = 3
		Z2 = X2 + Y2
	)
	const (
		X3 = `aaa
		sfafdas

		fdasfa`
	)
	/**
	const (
		A = iota
		B
		C
		D = 17
		E = iota
		F
	)
	**/
	const (
		A = 999
		B = iota
		C
	)

	println(A)
	println(B)
	println(C)
	const(
		朝　= 11
	)
}
