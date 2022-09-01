package main

import (
	"awesomeProject/HomeWork9/homework3.01/City"
	"awesomeProject/HomeWork9/homework3.01/Digit"
	"fmt"
	"math/rand"
)

/*
Создайте в проекте пакет(с произвольным названием),
который будет состоять из двух go файлов и содержать 2 функции,
по одной в каждом файле:
*City() - возвращает случайный город,
*Digit() - возвращает случайное число строчного типа
*Обе функции должны формировать результат с помощью функции Random
из пакета wordz. При этом не внося никаких изменение в пакет wordz.
выведите через fmt.Println результат функции City и Digit в файле main.go
*/
/*
С помощью утилиты go get, установите пакет для работы со строками xstrings.
В файле main.go примените фунцию Shaffle
из этого пакета к результату функции City().
И угадайте, какое название города вывелось
*/
func Shuffle(str string) string {
	if str == "" {
		return str
	}

	runes := []rune(str)
	index := 0

	for i := len(runes) - 1; i > 0; i-- {
		index = rand.Intn(i + 1)

		if i != index {
			runes[i], runes[index] = runes[index], runes[i]
		}
	}

	return string(runes)
}
func main() {
	fmt.Println("Random city:", Shuffle(City.City()))
	fmt.Println(Digit.Digit())

}
