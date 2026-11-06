package main

import (
	"fmt"
)

const basic int = 1
const floatingBall int = 2
const exit int = -1

func main() {
	var selectedTask int
	for selectedTask != -1 {
		fmt.Println("Выберите тип задачи:\n1) Тестовая\n2) Задача о погружении шара\n-1) Выход")
		_, err := fmt.Scan(&selectedTask)
		if err != nil || (selectedTask != basic && selectedTask != floatingBall && selectedTask != exit) {
			fmt.Println("Ошибка ввода")
			return
		}
		if selectedTask == basic {
			basicTask()
		}
		if selectedTask == floatingBall {
			floatingBallTask()
		}
	}
}
