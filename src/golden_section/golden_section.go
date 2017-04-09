package golden_section

import (
	"math"
	"utils"
)

var (
	/***input parameters***/
	eps   float64 = 0.01
	alpha float64 = (1 + math.Sqrt(5)) / 2

	a float64 = 0
	b float64 = 1
	/***end input parameters***/

	x float64 = 0
	y float64 = 0

	x1 float64 = 0
	x2 float64 = 0

	y1 float64 = 0
	y2 float64 = 0

	iterations int = 1
)

func GoldenRatio(polynomial []float64) (float64, float64, int) {
	iterations = 1
	x1 = b - (b-a)*(1/alpha)
	x2 = a + (b-a)*(1/alpha)
	y1 = utils.CalcPolynomialFunction(polynomial, x1)
	y2 = utils.CalcPolynomialFunction(polynomial, x2)
	for math.Abs(b-a) > eps {
		iterations++
		if y1 < y2 {
			b = x2
			x2 = x1
			x1 = b - (b-a)*(1/alpha)
			y2 = y1
			y1 = utils.CalcPolynomialFunction(polynomial, x1)
		} else {
			a = x1
			x1 = x2
			x2 = a + (b-a)*(1/alpha)
			y1 = y2
			y2 = utils.CalcPolynomialFunction(polynomial, x2)
		}
	}
	x = a
	y = y1
	return x, y, iterations
}
