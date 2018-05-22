package main

import (
	"./animal"
	"fmt"
)

func main() {
	fmt.Println(AppName())
	fmt.Println(animal.ElephantFeed())
	fmt.Println(animal.MonkeyFeed())
	fmt.Println(animal.RabbitFeed())
}
