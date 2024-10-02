package main

import (
	"AP1/tasks"
	"fmt"
)

func main() {
	var task string // Переменная для выбьора задания
	for {
		fmt.Print("\nВведите номер задания (1-6) или q чтобы выйти: ")
		fmt.Scan(&task) // Выбор желаемого задания

		switch task { // Выполнение задание в зависимости от введенного числа
		case "1":
			tasks.Task1()
		case "2":
			tasks.Task2()
		case "3":
			tasks.Task3()
		case "4":
			tasks.Task4()
		case "5":
			tasks.Task5()
		case "6":
			tasks.Task6()
		}

		if task == "q" { //Выход из цикла (завершение программы)
			break
		}
	}
}
