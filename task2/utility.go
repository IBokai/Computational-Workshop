package main

import (
	"fmt"
	"strings"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

const (
	ddColWidth = 14
	ddMaxOrder = 3
)

func dividedDifferenceTitle(order int) string {
	if order == 0 {
		return "f[z_i]"
	}
	return fmt.Sprintf("f[z_i..z_i+%d]", order)
}

func PrintDividedDifferences(nodes []Node, diff [][]float64) {
	orders := len(diff)
	cols := min(orders, ddMaxOrder)

	fmt.Println("Разделённые разности:")
	fmt.Printf("%*s", ddColWidth, "z_k")
	for j := 0; j < cols; j++ {
		fmt.Printf(" | %*s", ddColWidth, dividedDifferenceTitle(j))
	}
	fmt.Println()
	fmt.Println(strings.Repeat("-", ddColWidth+cols*(ddColWidth+3)))

	for i := 0; i < orders; i++ {
		fmt.Printf("%*.6f", ddColWidth, nodes[i].zk)
		for j := 0; j < cols; j++ {
			if j < len(diff[i]) {
				fmt.Printf(" | %*.6g", ddColWidth, diff[i][j])
			} else {
				fmt.Printf(" | %*s", ddColWidth, "")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func PrintNodes(nodes []Node) {
	fmt.Println("  z_k  |  f(z_k)  |  расстояние")
	fmt.Println("--------------------------------")
	for i, node := range nodes {
		fmt.Printf("%d) %6.4f  |  %6.4f   |  %6.4f\n", i+1, node.zk, node.fzk, node.distance)
	}
}