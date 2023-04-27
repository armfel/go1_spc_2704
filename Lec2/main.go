package main

import (
	"fmt"
	"math"
)

func main() {
	//Простейший вывод на консоль, println - вывод аргумента + '\n'
	fmt.Println("Hello world", "Hello another")
	fmt.Println("Second line")
	//Функция print - простой вывод аргумента
	fmt.Print("First")
	fmt.Print("Second")
	fmt.Println("Third")
	fmt.Println("Это моя вторая программа! Я рад, что учусь здесь!!")
	//Форматированный вывод
	fmt.Printf("Hello, my name is %s\nMy age is %d\n", "Bob", 42)
	///////////////////////////////////////////
	///////////////////////////////////////////
	//Декларирование переменных
	var age int
	fmt.Println("My age is:", age)
	age = 32
	fmt.Println("Age after assignment:", age)

	//Декларирование и инициализация пользовательским значением
	var height int = 185
	fmt.Println("My height is:", height)

	//В чем "полустрогость" типизации?
	var uid = 12345
	fmt.Println("My uid:", uid)
	fmt.Printf("%T\n", uid)

	//Декларирование и инициализация переменных одного типа (множественный случай)
	var firstVar, secondVar int = 20, 30
	fmt.Printf("FirstVar:%d SecondVar:%d\n", firstVar, secondVar)

	//Декларирование блока переменных
	var (
		personName string = "Bob"
		personAge  int    = 42
		personUID  int
	)

	fmt.Printf("Name: %s\nAge: %d\nUID: %d\n", personName, personAge, personUID)

	//Немного странного
	var a, b = 30, "Vova"
	fmt.Println(a, b)

	//Короткая декларация (короткое объявление)
	count := 10
	fmt.Println("Count:", count)
	count++
	fmt.Println("Count:", count)

	//Множественное присвание через :=
	aArg, bArg := 10, 30
	fmt.Println(aArg, bArg)
	aArg, bArg = 30, 40
	fmt.Println(aArg, bArg)

	//Исключения из правил
	bArg, cArg := 300, 400
	fmt.Println(bArg, cArg)

	//Пример
	width, length := 20.5, 30.2
	fmt.Printf("Min dimensional of rectangle is: %.3f\n", math.Min(width, length))

	//Еще пример
	var (
		word1 string = "имя"
		word2 string = "твое"
		word3 string = "мне"
		word4 string = "знакомо"
	)

	fmt.Print(word4, " ", word3, " ", word2, " ", word1, "\n", word3, " ", word1, " ", word4, " ", word2, "\n", word2, " ", word4, " ", word1, " ", word3, "\n")

}
