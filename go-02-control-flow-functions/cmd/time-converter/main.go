package main

import "fmt"

func main() {
	var totalMinutes int
	fmt.Print("Введите количество минут: ")

	_, err := fmt.Scan(&totalMinutes)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	if totalMinutes < 0 {
		fmt.Println("Ошибка: количество минут не должно быть отрицательным.")
		return
	}

	const min_in_hour = 60
	
	hours := totalMinutes / min_in_hour
	min_ost := totalMinutes % min_in_hour
	
	fmt.Printf("Колличество минут: %d\n", totalMinutes )
	fmt.Printf("Полных часов: %d\n", hours )
	fmt.Printf("Оставшихся минут: %d\n", min_ost)
}