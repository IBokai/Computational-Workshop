package main

import (
	"fmt"
	"math"
)

type Material struct {
	name string
	p    float64
}

func floatingBallTask() {
	fmt.Println("Выбрана задача о погружении шара")
	materials := []Material{
		{"Пробка", 0.25},
		{"Бамбук", 0.4},
		{"Сосна", 0.5},
		{"Кедр", 0.55},
		{"Дуб", 0.7},
		{"Бук", 0.75},
		{"Красное дерево", 0.8},
		{"Тиковое дерево", 0.85},
		{"Парафин", 0.9},
		{"Лёд", 0.92},
		{"Пчелиный воск", 0.95}}
	var r float64
	for {
		functions := []func(h float64) float64{func(h float64) float64 { return (math.Pi / 3) * (h*h*h - 3*h*h*r + 4*r*r*r*materials[0].p) },
			func(h float64) float64 { return (math.Pi / 3) * (h*h*h - 3*h*h*r + 4*r*r*r*materials[1].p) },
			func(h float64) float64 { return (math.Pi / 3) * (h*h*h - 3*h*h*r + 4*r*r*r*materials[2].p) },
			func(h float64) float64 { return (math.Pi / 3) * (h*h*h - 3*h*h*r + 4*r*r*r*materials[3].p) },
			func(h float64) float64 { return (math.Pi / 3) * (h*h*h - 3*h*h*r + 4*r*r*r*materials[4].p) },
			func(h float64) float64 { return (math.Pi / 3) * (h*h*h - 3*h*h*r + 4*r*r*r*materials[5].p) },
			func(h float64) float64 { return (math.Pi / 3) * (h*h*h - 3*h*h*r + 4*r*r*r*materials[6].p) },
			func(h float64) float64 { return (math.Pi / 3) * (h*h*h - 3*h*h*r + 4*r*r*r*materials[7].p) },
			func(h float64) float64 { return (math.Pi / 3) * (h*h*h - 3*h*h*r + 4*r*r*r*materials[8].p) },
			func(h float64) float64 { return (math.Pi / 3) * (h*h*h - 3*h*h*r + 4*r*r*r*materials[9].p) },
			func(h float64) float64 { return (math.Pi / 3) * (h*h*h - 3*h*h*r + 4*r*r*r*materials[10].p) }}
		fmt.Println("Введите радиус шара(в метрах) (0 для выхода):")
		fmt.Print("r=")
		fmt.Scan(&r)
		if r == 0 {
			return
		}
		if r < 0 {
			fmt.Println("Радиус шара должен быть положительным!")
			continue
		}
		for i, f := range functions {
			intervals, _ := rootSeparation(f, 0, 2*r, 1000)
			h, _, _, _ := bisectionMethod(intervals[0].A, intervals[0].B, f, -5)
			fmt.Print(i+1, ")", materials[i].name, " --- ")
			fmt.Printf("%.3f\n", h)
		}
	}
}
