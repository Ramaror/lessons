package main

import "fmt"

func main() {
	type AmericanVelocity float64
	type EuropeanVelocity int

	var a AmericanVelocity
	var b EuropeanVelocity
	var c float64
	var d float64

	c = 130.0 * 1.609
	a = AmericanVelocity(d)
	fmt.Println(a)

	d = 120.4 * 3.6
	b = EuropeanVelocity(c)
	fmt.Println(b)

}
