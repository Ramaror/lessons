package main

import "fmt"

func main() {
	day := []string{"понедельник", "вторник", "среда", "четверг", "пятница", "суббота", "воскресенье"}
	rab := day[:5]
	vih := day[5:7]
	fmt.Println("Неделя:", day)
	fmt.Println("Рабочие дни:", rab)
	fmt.Println("Выходные дни:", vih)
	rab = append(rab, vih...)
	fmt.Println("Обьъединение рабочих и выходных дней:", rab)
}
