package main

func rootSeparation(f MathFunctionType, a_border, b_border float64, n int) ([]Interval, error) {
	var intervals []Interval
	h := (b_border - a_border) / float64(n)
	a_temp_border := a_border
	b_temp_border := a_temp_border + h
	for b_temp_border <= b_border {
		if f(a_temp_border)*f(b_temp_border) <= 0 {
			intervals = append(intervals, Interval{a_temp_border, b_temp_border})
		}
		a_temp_border = b_temp_border
		b_temp_border = a_temp_border + h
	}
	return intervals, nil
}
