package main

import "fmt"

// statusLabel превращает код статуса в понятную человеку подпись.
func statusLabel(status string) string {
	var status_translate string

	switch status {
	case "new":
		status_translate = "Новая"

	case "in_progress":
		status_translate = "В работе"

	case "done":
		status_translate = "Завершена"

	default:
		status_translate = "Неизвестный статус"
	}

	// TODO: обработайте new, in_progress, done через switch.
	// TODO: предусмотрите ветку default для незнакомого статуса.
	return status_translate
}

// countStatus считает количество элементов с заданным статусом.
func countStatus(statuses []string, target string) int {
	counter := 0

	for _, status := range statuses {
		if status == target {
			counter++
		}
	}

	// TODO: пройдите по срезу через for range.
	// TODO: сравните каждый элемент с target и верните счётчик.
	return counter
}

func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	statuses = append(statuses, "new")
	// TODO: добавьте ещё один статус new через append.

	fmt.Printf("Задач всего: %d\n", len(statuses))
	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}

	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))

	// TODO: создайте map[string]int для подсчёта статусов.
	// TODO: выведите new, in_progress, done в фиксированном порядке.

	counts := make(map[string]int)

	for _, status := range statuses {
		counts[status]++
	}

	fmt.Printf("Статистика: new=%d, in_progress=%d, done=%d",
		counts["new"],
		counts["in_progress"],
		counts["done"],
	)

}
