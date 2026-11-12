package main

import "fmt"

type Task struct {
	ID       int
	Title    string
	Status   string
	Priority int
}

func (t Task) Summary() string {
	// TODO: верните "#ID Название [статус], приоритет N".
	return fmt.Sprintf("#%d %s [%s], приоритет %d",
		t.ID, t.Title, t.Status, t.Priority)
}

func (t *Task) Complete() {
	// TODO: присвойте статус done исходной задаче.
	t.Status = "done"
}

func (t *Task) SetPriority(priority int) bool {
	// TODO: проверьте диапазон 1–3, при ошибке объект не менять.

	if priority >= 1 && priority <= 3 {
		t.Priority = priority
	} else {
		return false
	}

	return true
}

func findTaskByID(tasks []Task, id int) *Task {
	// TODO: верните &tasks[i] при совпадении, иначе nil.
	for i, task := range tasks {
		if task.ID == id {
			return &tasks[i]
		}
	}

	return nil
}

func tasksByStatus(tasks []Task, status string) []Task {
	var new_tasks []Task

	for _, task := range tasks {
		if task.Status == status {
			new_tasks = append(new_tasks, task)
		}
	}

	return new_tasks //вернет nil если пустой
}

func main() {
	tasks := []Task{
		{ID: 1, Title: "Установить Go", Status: "done", Priority: 1},
		{ID: 2, Title: "Освоить функции", Status: "in_progress", Priority: 2},
		{ID: 3, Title: "Изучить структуры", Status: "new", Priority: 3},
	}

	// TODO: найдите ID=2, завершите и установите приоритет 3.

	for i, task := range tasks {
		if task.ID == 2 {
			tasks[i].Priority = 3
			tasks[i].Status = "done"
		}
	}

	// // TODO: посчитайте завершённые задачи.

	counter := 0
	tasks[0].Complete()

	for _, task := range tasks {
		fmt.Println(task.Summary())
		if task.Status == "done" {
			counter++
		}
	}

	fmt.Printf("Задач завершено: %d\n", counter)

	if findTaskByID(tasks, 99) == nil {
		fmt.Println("Задача не найдена")
	}
}
