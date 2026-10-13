package main

import "fmt"

func main() {
	const language = "Go"
	var name string = "Даниил"
	completed := 0
	var ready bool = true
	var hours float64 = 1

	fmt.Println("Изучаем", language)
	fmt.Printf("Участник: %s\n", name)
	fmt.Printf("Выполнено: %d\n", completed)
	fmt.Printf("Готов продолжать: %t\n", ready)
	fmt.Printf("Продолжительность: %.1f часа\n", hours)

	completed = 3
	fmt.Printf("После изменения: %d\n", completed)
}
