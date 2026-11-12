package main

import "fmt"

func main() {
	const language = "Бэкэнд разработка на языке Go"
	name := "Савелий"
	completed := 1
	var group = "КЕФТЕМЕ-02-67"
	var ready = true
	

	fmt.Printf("Участник: %s\n", name)
	fmt.Printf("Группа: %s\n", group)
	fmt.Printf("Курс: %s\n", language)
	fmt.Printf("Завершено работ: %d\n", completed)
	fmt.Printf("Интерес к бюкэнду: %t\n", ready)
}