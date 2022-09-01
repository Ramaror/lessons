package main

import (
	newcolor "awesomeProject/HomeWork9/main/color"
	"awesomeProject/HomeWork9/main/wordz"
	"fmt"
	"github.com/fatih/color"
)

func main() {
	newcolor.Greet()
	fmt.Println("Hello world!")
	color.Red("Hello World again")

	for i := 0; i < 5; i++ {
		fmt.Println(wordz.Hello)
		fmt.Println(wordz.Random())
	}
}
