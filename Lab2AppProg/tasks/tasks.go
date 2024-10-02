package tasks

import (
	"fmt"
	"unicode/utf8"
)

func Task1() {
	var number, check int
	fmt.Print("Введите целое число: ")
	fmt.Scan(&number)
	check = number % 2 // Определеяет остаток от деления на число

	if check == 0 {
		fmt.Printf("Число %d четное", number)
	} else {
		fmt.Printf("Число %d нечетное", number)
	}
}

func claccifyNumber(num float32) string {
	if num > 0 { // Условные выражения для проверки на знак числа
		return "Positive"
	} else if num < 0 {
		return "Negative"
	} else {
		return "Zero"
	}
}

func Task2() {
	var num float32
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	result := claccifyNumber(num)
	fmt.Printf("Ваше число %s", result)
}

func Task3() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	fmt.Printf("Все числа от 1 до %d:", number)
	for i := 1; i <= number; i++ { //Вывод числа с помощью цикла for
		fmt.Printf("\n%d", i)
	}
}

func lengthString(str string) int {
	return utf8.RuneCountInString(str) // Вычисления длины строки с помощью функции RuneCountInString (использовал вместо len так как она учитывает все символы unicode)
}

func Task4() {
	var input string
	fmt.Print("Введите строку разделяя слова точками: ")
	fmt.Scan(&input)
	result := lengthString(input)
	fmt.Printf("Длина вашей строки: %d", result)
}

type Rectangle struct { // Структура принимающая 2 значения
	widith float32
	height float32
}

func calculateArea(r Rectangle) float32 {
	return r.widith * r.height
}

func Task5() {
	var w, h float32
	fmt.Print("Введите ширину прямоугольника: ")
	fmt.Scan(&w)
	fmt.Print("Введите высоту прямоугольника: ")
	fmt.Scan(&h)

	rect := Rectangle{ // Создание экземпляра структуры и ввод данных
		widith: w,
		height: h,
	}

	area := calculateArea(rect)
	fmt.Printf("Площадь введенного прямоугольника: %f условных единиц", area)
}

func avg(n1 int, n2 int) float32 {
	return float32(n1+n2) / 2 // Приведения int результата к float32
}

func Task6() {
	var num1, num2 int
	fmt.Print("Введите первое целое число: ")
	fmt.Scan(&num1)
	fmt.Print("Введите второе целое число: ")
	fmt.Scan(&num2)

	result := avg(num1, num2)
	fmt.Print(result)
}
