package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

/*
В файле main.go напишите код,
который считывает данные из файла int.txt
и построчно записывает их в файл out.txt, нумеруя каждую строку.
Если файла out.txt нет, то он должен создаваться.
С помощью отложенных вызово закройте файловые дескрипторы. При закрытии файла out.txt
программа должна вывести в консоль, сколько строк и байт было записано в файл.
*/
func HomeWork() {
	file, err := os.Open("C:\\Users\\doksh\\GolandProjects\\awesomeProject\\HomeWork8\\data\\in")
	qw, err := os.Create("C:\\Users\\doksh\\GolandProjects\\awesomeProject\\HomeWork8\\data\\out.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	defer qw.Close()

	reader := bufio.NewReader(file)
	i := 0
	var dot string = ". "
	for {
		i++
		line, err := reader.ReadString('\n')

		qw.WriteString(strconv.Itoa(i))
		qw.WriteString(dot)
		qw.WriteString(line)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
	}

	fmt.Println("Записано", i, "строк")
}
func main() {
	//Реализуйте функцию logTIme()
	start := time.Now()
	HomeWork()
	duration := time.Since(start)
	fmt.Printf("Elapsed time: %v", duration)
}
