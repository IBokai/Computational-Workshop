package main

func NewtonInterpolation(nodes []Node, n int, x float64) float64 {
	diffs := make([][]float64, n+1)
	for i := range diffs {
		diffs[i] = make([]float64, n+1-i)
	}
	for i := 0; i <= n; i++ {
		diffs[i][0] = nodes[i].fzk
	}
	for j := 1; j <= n; j++ {
		for i := 0; i <= n-j; i++ {
			numerator := diffs[i+1][j-1] - diffs[i][j-1]
			denominator := nodes[i+j].zk - nodes[i].zk
			diffs[i][j] = numerator / denominator
		}
	}
	//PrintDividedDifferences(nodes, diffs)

	result := diffs[0][0]
	for k := 1; k <= n; k++ {
		product := 1.0
		for j := 0; j < k; j++ {
			product *= (x - nodes[j].zk)
		}
		result += diffs[0][k] * product
	}
	return result
}
