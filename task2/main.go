package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	for {
		clearScreen()
		fmt.Println("Задача алгебраического интерполирования")
		fmt.Println("Выберите функцию из списка:")
		fmt.Println("1) e^(-x) - x^2/2")
		fmt.Println("2) x^2/(1+x^2)")
		fmt.Println("3) 2*sin(x) - x/2")
		fmt.Println("4) x^3 - 4*x + 2")
		fmt.Println("0) Выход")
		functions := []func(float64) float64{
			func(x float64) float64 { return math.Exp(-x) - (x*x)/2 },
			func(x float64) float64 { return (x * x) / (1 + x*x) },
			func(x float64) float64 { return 2*math.Sin(x) - x/2 },
			func(x float64) float64 { return x*x*x - 4*x + 2 }}
		var selectedFunctionN int
		_, err := fmt.Scan(&selectedFunctionN)
		if err != nil || selectedFunctionN < 0 || selectedFunctionN > len(functions) {
			fmt.Println("Ошибка выбор функции")
			continue
		}
		if selectedFunctionN == 0 {
			break
		}

		selectedFunction := functions[selectedFunctionN-1]
		fmt.Println("Введите число значений в таблице:")
		var m int
		_, err = fmt.Scan(&m)
		for err != nil || m < 1 {
			fmt.Println("Ошибка ввода числа значений в таблице")
			fmt.Println("Введите число значений в таблице:")
			_, err = fmt.Scan(&m)
		}

		fmt.Println("Введите [A, B]: ")
		var a, b float64
		_, err = fmt.Scan(&a, &b)
		for err != nil || a > b {
			fmt.Println("Ошибка ввода [A, B]")
			fmt.Println("Введите [A, B]: ")
			_, err = fmt.Scan(&a, &b)
		}

		fmt.Println("Выберите характер узлов: ")
		fmt.Println("1) Равностоящие")
		fmt.Println("2) Случайные")
		fmt.Println("3) Ввести самому")
		var rootTypeN int
		_, err = fmt.Scan(&rootTypeN)
		for err != nil || rootTypeN < 1 || rootTypeN > 3 {
			fmt.Println("Ошибка выбора характера узлов")
			fmt.Println("Выберите характер узлов: ")
			fmt.Println("1) Равностоящие")
			fmt.Println("2) Случайные")
			fmt.Println("3) Ввести самому")
			_, err = fmt.Scan(&rootTypeN)
		}
		var nodes []Node
		switch rootTypeN {
		case 1:
			nodes = GetEquidistantNodes(m, a, b, selectedFunction)
		case 2:
			nodes = GetRandomNodes(m, a, b, selectedFunction)
		case 3:
			nodes = GetUserNodes(m, a, b, selectedFunction)
		}
		PrintNodes(nodes)

		var x float64
		var ok string = "y"

		for strings.ToLower(ok) == "y" {
			fmt.Println("Введите точку интерполирования x: ")
			_, err = fmt.Scan(&x)
			for err != nil {
				fmt.Println("Ошибка ввода точки интерполирования")
				fmt.Println("Введите точку интерполирования x: ")
				_, err = fmt.Scan(&x)
			}

			for i := 0; i < len(nodes); i++ {
				nodes[i].distance = math.Abs(nodes[i].zk - x)
			}
			sort.Slice(nodes, func(i, j int) bool {
				return nodes[i].distance < nodes[j].distance
			})

			var n int
			fmt.Printf("Введите степень интерполяционного многочлена n(<=%d)\n", m-1)
			_, err = fmt.Scan(&n)
			for err != nil || n > m-1 || n < 1 {
				fmt.Println("Ошибка ввода степени интерполяционного многочлена")
				fmt.Println("Введите степень интерполяционного многочлена n(<=m)")
				_, err = fmt.Scan(&n)
			}
			PrintNodes(nodes[:n+1])
			lagrangeValue := LagrangeInterpolation(nodes, n, x)
			newtonValue := NewtonInterpolation(nodes, n, x)
			fmt.Println("Значение по Лагранжу:", lagrangeValue)
			fmt.Println("Абсолютная погрешность по Лагранжу:", math.Abs(selectedFunction(x)-lagrangeValue))
			fmt.Println("Значение по Ньютону:", newtonValue)
			fmt.Println("Абсолютная погрешность по Ньютону:", math.Abs(selectedFunction(x)-newtonValue))
			fmt.Println("\nРазность:", math.Abs(lagrangeValue-newtonValue))
			fmt.Println("Хотите ввести другой x? (y/n)")
			_, err := fmt.Scan(&ok)
			if err != nil {
				fmt.Println("Fatal error")
				return
			}
		}
	}
}
