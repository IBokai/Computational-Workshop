package main

import (
	"fmt"
	"math"
)

func LagrangeInterpolation(nodes []Node, n int, x float64) float64 {
	pLagrange := 0.0
	coeffsSum := 0.0
	lebesgue := 0.0
	for k := 0; k <= n; k++ {
		lK := 1.0
		for j := 0; j <= n; j++ {
			if j != k {
				numerator := x - nodes[j].zk
				denominator := nodes[k].zk - nodes[j].zk
				lK *= numerator / denominator
			}
		}
		pLagrange += nodes[k].fzk * lK
		coeffsSum += lK
		lebesgue += math.Abs(lK)
	}
	fmt.Println("Сумма коэффицентов: ", coeffsSum)
	fmt.Println("Значение функции Лебега: ", lebesgue)
	return pLagrange
}
