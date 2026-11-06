package main

import (
	"fmt"
	"math"
)

func basicTask() {
	fmt.Println("Выбрана задача о нахождении корней нечётной кратности нелинейного уравнения")
	fmt.Println("Выберите функцию из списка:")
	functions := []MathFunction{
		{func(x float64) float64 { return x*math.Sin(x) - 1 }, func(x float64) float64 { return math.Sin(x) + x*math.Cos(x) }, "x*sin(x) - 1"},
		{func(x float64) float64 { return 8*math.Cos(x) - x - 6 }, func(x float64) float64 { return -8*math.Sin(x) - 1 }, "8*cos(x) - x - 6"},
		{func(x float64) float64 { return 10*math.Cos(x) - 0.1*x*x }, func(x float64) float64 { return -10*math.Sin(x) - x/5 }, "10*cos(x) - 0.1*x^2"},
		{func(x float64) float64 { return 4*math.Cos(x) + 0.3*x }, func(x float64) float64 { return -4*math.Sin(x) + 0.3 }, "4*cos(x) + 0.3*x"},
		{func(x float64) float64 { return 5*math.Sin(2*x) - math.Sqrt(1-x) }, func(x float64) float64 { return 10*math.Cos(2*x) + 1/(2*math.Sqrt(1-x)) }, "5*sin(2x) - sqrt(1-x)"}}
	for i, f := range functions {
		fmt.Print(i+1, ") ")
		fmt.Println(f.name)
	}
	var selectedFunction int
	_, err := fmt.Scan(&selectedFunction)
	if err != nil || selectedFunction < 1 || selectedFunction > 5 {
		fmt.Println("Ошибка выбора функции")
		return
	}
	clearScreen()
	fmt.Println("Выбрана функция ", functions[selectedFunction-1].name)
	var a_border, b_border float64
	fmt.Println("Введите через пробел концы отрезка, на котором хотите найти корни:")
	_, err = fmt.Scan(&a_border, &b_border)
	var n int
	confirmation := "y"
	var intervals []Interval

	for confirmation == "y" {
		fmt.Println("Введите число N разбиений отрезка")
		if _, err = fmt.Scan(&n); err != nil || n <= 2 {
			fmt.Println("Ошибка ввода числа разбиений отрезка")
			return
		}
		intervals, err = rootSeparation(functions[selectedFunction-1].function, a_border, b_border, n)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Было найдено", len(intervals), "отрезков перемены знака")
		for i, interval := range intervals {
			fmt.Println(i+1, ")", interval)
		}
		fmt.Println("Хотите ввести другое число разбиений отрезка?(y/n)")
		_, err = fmt.Scan(&confirmation)
		if err != nil {
			fmt.Println("Ошибка подтверждения выбора нового числа разбиений отрезка")
			return
		}
	}

	var selectedInterval int
	confirmation = "y"
	for confirmation == "y" {
		fmt.Println("Выберите интервал, в котором хотите уточнять корень:")
		_, err := fmt.Scan(&selectedInterval)
		if err != nil || selectedInterval > len(intervals) || selectedInterval <= 0 {
			fmt.Println("Ошибка выбора интервала")
			continue
		}
		selectedInterval = selectedInterval - 1
		fmt.Println("Выбран интервал: ", intervals[selectedInterval])

		var epsilon_power int
		epsilon_confirmation := "y"
		for epsilon_confirmation == "y" {
			fmt.Println("Выберите точность уточнения корня:")
			fmt.Print("10^-")
			if _, err = fmt.Scan(&epsilon_power); err != nil {
				fmt.Println("Ошибка ввода точности")
				continue
			}
			fmt.Printf("%-25s | %-16s | %-22s | %-15s\n", "Название метода", "Количество шагов", "Приближенный корень", "Модуль невязки")

			x, steps, magnitude, _ := bisectionMethod(intervals[selectedInterval].A, intervals[selectedInterval].B, functions[selectedFunction-1].function, -epsilon_power)
			fmt.Printf("%-25s | %-16d | %-22.15f | %-15.2e\n", "Метод бисекции", steps, x, magnitude)

			x, steps, magnitude, _ = newtonsMethod(intervals[selectedInterval].A, intervals[selectedInterval].B, functions[selectedFunction-1].function, functions[selectedFunction-1].derivative, -epsilon_power)
			fmt.Printf("%-25s | %-16d | %-22.15f | %-15.2e\n", "Метод Ньютона", steps, x, magnitude)

			x, steps, magnitude, _ = modNewtonsMethod(intervals[selectedInterval].A, intervals[selectedInterval].B, functions[selectedFunction-1].function, functions[selectedFunction-1].derivative, -epsilon_power)
			fmt.Printf("%-25s | %-16d | %-22.15f | %-15.2e\n", "Модифициц. Мет. Ньютона", steps, x, magnitude)

			x, steps, magnitude, _ = secantMethod(intervals[selectedInterval].A, intervals[selectedInterval].B, functions[selectedFunction-1].function, -epsilon_power)
			fmt.Printf("%-25s | %-16d | %-22.15f | %-15.2e\n", "Метод секущих", steps, x, magnitude)


			fmt.Println("Хотите изменить точность уточнения корня?(y/n)")
			_, err = fmt.Scan(&epsilon_confirmation)
			if err != nil {
				fmt.Println("Ошибка выбора новой точности уточнения корня")
				return
			}
		}
		fmt.Println("Хотите выбрать другой интервал?(y/n)")
		_, err = fmt.Scan(&confirmation)
		if err != nil {
			fmt.Println("Ошибка подтверждения выбора нового интервала")
			return
		}
	}
}
