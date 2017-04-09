package utils

import "math"

func CalcPolynomialFunction(polynomial []float64, x float64) float64 {
	var res float64
	for i := 0; i < len(polynomial); i++ {
		res += polynomial[i] * math.Pow(x, float64(i))
	}
	return res
}
