package main

import "fmt"

func main() {
	var A *int
	var B int
	A = &B
	fmt.Println(*A)

	B = *A
	fmt.Println(B)
}
