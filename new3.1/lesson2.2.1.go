package main

import (
	"fmt"
)

func main() {
	var R *float64
	var sum float64
	const p = 3.14
	sum = 35 / (2 * p)
	R = &sum
	ot := fmt.Sprintf("%.2f", *R)
	fmt.Println("Радиус окружности:", ot)
	//fmt.Println("Радиус оркужности:", math.Round(*R))
	S := p * (*R * *R)
	ot2 := fmt.Sprintf("%.2f", S)
	fmt.Println("Площадь круга:", ot2)
	//fmt.Println("Площадь круга:", math.Round(S))
}
