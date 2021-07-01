package main

import (
	"fmt"

	"hoge/animals"
)

func main() {
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())

	fmt.Println("hoge")
	fmt.Println("fuga")
	fmt.Println("bar")

	fmt.Println(AppName())

}
