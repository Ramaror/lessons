package main

import "fmt"

func contains(slice []string, x string) bool {
	for _, value := range slice {
		if value == x {
			return true
		}
	}
	return false
}

func getMax(num ...int) int {
	max := num[0]
	for _, value := range num {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	fmt.Println(contains([]string{"a", "b", "c"}, "b"))
	fmt.Println(getMax(1, 2, 4, 2, 3))
}
