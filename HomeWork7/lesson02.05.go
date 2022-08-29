package main

import "fmt"

// Напишите функцию contains, котоаря принимает на вход два параметра: слайс строк а
// и строку х. Функция должна проверять, содержится ли строка х в слайсе а, и возвращать будево значение
func contains(slice []string, x string) bool {
	for _, value := range slice {
		if value == x {
			return true
		}
	}
	return false
}

// Создайте вариативную функцию getMax, которая находит максимальное целое число и переданных на вход параметров
func getMax(num ...int) int {
	max := num[0]
	for _, value := range num {
		if value > max {
			max = value
		}
	}
	return max
}

// Выведите на экран рузельтат вызова функций
func main() {
	fmt.Println(contains([]string{"a", "b", "c"}, "b"))
	fmt.Println(getMax(1, 2, 4, 2, 3))
}
