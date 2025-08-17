package main

import (
	"fmt"
	"main/internal/domain"
)

func main() {
	list := domain.NewTaskList()
	task1, err := domain.NewTask("Задача 1", "Помыть попу")
	if err != nil {
		fmt.Println("error task creating: ", err)
	}
	list.CreateTask(task1)
	for _, v := range list.Tasks {
		fmt.Println(v.Description)
	}
}
