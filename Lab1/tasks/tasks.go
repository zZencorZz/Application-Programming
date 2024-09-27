package tasks

import (
	"fmt"
	"time"
)

func Task1() {
	for {
		current_time := time.Now().Format("2006-01-02 15:04:05") /* Передаем в переменную текущее время
		и форматируем его с помощью спец значения */
		fmt.Println("\033[H\033[2J") // Команда для очистки терминала
		fmt.Println(current_time)    // Выводим форматированное время
		time.Sleep(1 * time.Second) // Задержка для более приятного вывода
	}
}

func Task2() {
	// Задаем переменные различных типов
	var var_int int = 10
	var var_float64 float64 = 10.10
	var var_str string = "ten"
	var var_bool bool = true

	/* Помещаем переменные в словарь Map( позволит добавить к каждому элементу массива ключ)
	с типом пустой интерфейс Interface (позвоялет заполнить массив переменными различных типов) */
	var mix_map = map[string]interface{}{
		"int":     var_int,
		"float64": var_float64,
		"string":  var_str,
		"bool":    var_bool,
	}

	for key, value := range mix_map { //Используем цикл для прохождения по элементам массива
		fmt.Printf("%s: %v\n", key, value) /*выводим значения с помощью форматированной строки
		(%s/v форматные спецификаторы string и variable)*/
	}
}

func Task3() {
	/* Задаем переменные различных типов с помощью
	краткой записи без объявления типов */
	var_int := 10
	var_float64 := 10.10
	var_str := "ten"
	var_bool := true

	var mix_map = map[string]interface{}{
		"int":     var_int,
		"float64": var_float64,
		"string":  var_str,
		"bool":    var_bool,
	}

	for key, value := range mix_map { //Используем цикл для прохождения по элементам массива
		fmt.Printf("%s: %v\n", key, value)
	}
}

func Task4() {
	var in [2]int       // Массив для целых чисел
	var operator string // Переменнаядля выбор оператора

	for i := 0; i < 2; i++ { // Ввод чисел в цикле для оптимизации кода
		fmt.Printf("Введите целое число int%d: ", i+1)
		fmt.Scan(&in[i])
	}

	fmt.Print("Введите желаемую операцию (+, -, *, /): ")
	fmt.Scan(&operator) // Выбор операции

	switch operator { // Варианты выполнения операций
	case "+":
		fmt.Printf("%d + %d = %d", in[0], in[1], in[0]+in[1])

	case "-":
		fmt.Printf("%d - %d = %d", in[0], in[1], in[0]-in[1])

	case "*":
		fmt.Printf("%d * %d = %d", in[0], in[1], in[0]*in[1])

	case "/":
		fmt.Printf("%d / %d = %d", in[0], in[1], in[0]/in[1])

	default:
		fmt.Print("Ошибка: неверная операция")
	}
}

func Task5() {
	var in [2]float32   // Массив для дробных чисел
	var operator string // Переменнаядля выбор оператора

	for i := 0; i < 2; i++ { // Ввод чисел в цикле для оптимизации кода
		fmt.Printf("Введите дробное число float%d: ", i+1)
		fmt.Scan(&in[i])
	}

	fmt.Print("Введите желаемую операцию (+, -): ")
	fmt.Scan(&operator) // Выбор операции

	switch operator { // Варианты выполнения операций
	case "+":
		fmt.Printf("%f + %f = %f", in[0], in[1], in[0]+in[1])

	case "-":
		fmt.Printf("%f - %f = %f", in[0], in[1], in[0]-in[1])

	default:
		fmt.Print("Ошибка: неверная операция")
	}
}

func Task6() {
	var num [3]float32 // Массив для введенных чисел

	for i := 0; i < 3; i++ { //ввод чисел в цикле для оптимизации кода
		fmt.Printf("Введите число%d: ", i+1)
		fmt.Scan(&num[i])
	}

	avg_number := (num[0] + num[1] + num[2]) / 3 // Вычисление ср. арифм.
	fmt.Printf("Среднее арифметическое ваших чисел: %f", avg_number)
}
