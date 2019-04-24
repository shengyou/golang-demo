package main

import (
	"fmt"
	"golang-demo/greeting"
	"golang-demo/repeat"
)

func main() {
	fmt.Println(greeting.Greeting("Shengyou", "Mandarin"))
	fmt.Println(repeat.Repeat(3, "b"))
}
