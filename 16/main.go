package main

import (
	_ "math"
)

func main() {
	//println(AorB()) //A
	for {
		//無限ループ
		//処理
		continue
	}
	for i := 0; i < 10; i++ {

	}
	x := 0
	if x == 1 {

	}
	if true {

	}
	/**
	if (1){

	}
	**/
	/**
	if x, err := 1, 2; err != nil {
		//エラー処理
	}
	**/
	z := 5
	if _, err := doSomething(); err != nil {
		//エラーハンドリング
	}

}

func AorB() (b string) {
	b = "A"
	{
		b := "B"
		print(b) //B
	}
	return
}
