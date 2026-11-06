package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

type Node struct {
	zk       float64
	fzk      float64
	distance float64
}

func GetEquidistantNodes(m int, a, b float64, f func(float64) float64) []Node {
	result := make([]Node, m)
	for i := range m {
		zk := a + ((b-a)/float64(m-1))*float64(i)
		result[i] = Node{zk, f(zk), -1}
	}
	return result
}

func GetRandomNodes(m int, a, b float64, f func(float64) float64) []Node {
	result := make([]Node, m)
	for i := 0; i < m; {
		zk := a + rand.Float64()*(b-a)
		if slices.ContainsFunc(result[:i], func(r Node) bool { return r.zk == zk }) {
			continue
		}
		result[i] = Node{zk, f(zk), -1}
		i++
	}
	return result
}

func GetUserNodes(m int, a, b float64, f func(float64) float64) []Node {
	result := make([]Node, m)
	fmt.Println("Введите узлы интерполяции: ")
	var zk float64
	for i := 0; i < m; {
		fmt.Printf("%d: \n", i+1)
		_, err := fmt.Scan(&zk)
		if err != nil {
			fmt.Println("Недопустимый узел")
			continue
		}
		if zk < a || zk > b {
			fmt.Println("Введён узел за пределами [A, B]")
			continue
		}
		if slices.ContainsFunc(result[:i], func(r Node) bool { return r.zk == zk }) {
			fmt.Println("Такой узел уже введён")
			continue
		}
		result[i] = Node{zk, f(zk), -1}
		i++
	}

	return result
}
