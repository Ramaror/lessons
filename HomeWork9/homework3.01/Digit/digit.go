package Digit

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var Prefix = "Random digit: "
var Words = []string{"One", "Two", "Three", "Four", "Five"}

func init() {
	fmt.Println("Function init in package digit")
}

func Digit() string {
	max := len(Words)
	var r, _ = rand.Int(rand.Reader, big.NewInt(int64(max)))
	return get(r.Int64())
}
func get(index int64) string {
	return Prefix + Words[index]
}
