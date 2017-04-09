package newton

import (
	"math"
	"utils"
)

var (
	/***input parameters***/
	eps float64 = 0.01
	x0  float64 = 0
	/***end input parameters***/

	x float64 = 0
	y float64 = 0

	x1 float64 = 0

	y1         float64 = 0
	y2         float64 = 0
	iterations int     = 1
)

func Newton(polynomial []float64, derivative1 []float64, derivative2 []float64) (float64, float64, int) {
	iterations = 1
	y1 = utils.CalcPolynomialFunction(derivative1, x0)
	y2 = utils.CalcPolynomialFunction(derivative2, x0)
	x1 = x0 - y1/y2
	for math.Abs(x1-x0) > eps {
		iterations++
		x0 = x1
		y1 = utils.CalcPolynomialFunction(derivative1, x0)
		y2 = utils.CalcPolynomialFunction(derivative2, x0)
		x1 = x0 - y1/y2
	}
	x = x1
	y = utils.CalcPolynomialFunction(polynomial, x)
	return x, y, iterations
}
