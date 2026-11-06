package main

import "math"

func bisectionMethod(a, b float64, f MathFunctionType, epsilon int) (float64, int, float64, error) {
	c := (a + b) / 2
	steps := 0
	for b-a > 2*math.Pow(10, float64(epsilon)) {
		c = (a + b) / 2
		if f(a)*f(c) <= 0 {
			b = c
		} else {
			a = c
		}
		steps++
	}
	x := (a + b) / 2
	return x, steps, math.Abs(f(x) - 0), nil
}

func newtonsMethod(a, b float64, f, df MathFunctionType, epsilon int) (float64, int, float64, error) {
	x0 := (a + b) / 2
	x1 := x0 - f(x0)/df(x0)
	steps := 1
	for math.Abs(x1-x0) > math.Pow(10, float64(epsilon)) {
		x0 = x1
		x1 = x0 - f(x0)/df(x0)
		steps++
	}
	return x1, steps, math.Abs(f(x1) - 0), nil
}

func modNewtonsMethod(a, b float64, f, df MathFunctionType, epsilon int) (float64, int, float64, error) {
	x0 := (a + b) / 2
	c := df(x0)
	x1 := x0 - f(x0)/c
	steps := 1
	for math.Abs(x1-x0) > math.Pow(10, float64(epsilon)) {
		x0 = x1
		x1 = x0 - f(x0)/c
		steps++
	}
	return x1, steps, math.Abs(f(x1) - 0), nil
}

func secantMethod(a, b float64, f MathFunctionType, epsilon int) (float64, int, float64, error) {
	x0 := a
	x1 := b
	x2 := x1 - (f(x1)/(f(x1)-f(x0)))*(x1-x0)
	steps := 1
	for math.Abs(x2-x1) > math.Pow(10, float64(epsilon)) {
		x0 = x1
		x1 = x2
		x2 = x1 - (f(x1)/(f(x1)-f(x0)))*(x1-x0)
		steps++
	}
	return x2, steps, math.Abs(f(x2) - 0), nil
}

