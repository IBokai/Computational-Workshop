package main

type MathFunctionType = func(float64) float64

type MathFunction struct {
	function   MathFunctionType
	derivative MathFunctionType
	name       string
}

type Interval struct {
	A float64
	B float64
}
