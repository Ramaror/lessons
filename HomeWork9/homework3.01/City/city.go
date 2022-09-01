package City

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var Citys = []string{"Grozniy", "Argun", "Gudermes", "Urus-martan", "Yandi"}

func init() {
	fmt.Println("Function init in package city")
}

func City() string {
	max := len(Citys)
	var r, _ = rand.Int(rand.Reader, big.NewInt(int64(max)))
	return get(r.Int64())
}
func get(index int64) string {
	return Citys[index]
}
