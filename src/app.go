package main

import (
	"fmt"
	"golden_ratio"
	"newton"
)

var polynomial = []float64{0.589431851, -1.552247823, 1.57128111, -0.012720386, -0.106634612, 0.001195152}

var derivative1 = []float64{-1.552247823, 3.14256222, -0.038161158, -0.426538448, 0.00597576}

var derivative2 = []float64{3.14256222, -0.076322316, -1.279615344, 0.02390304}

func main() {
	var x, y, iterations = golden_ratio.GoldenRatio(polynomial)
	fmt.Println("Метод Золотого Сечения: x=", x, " y=", y, " Итераций=", iterations)

	var x1, y1, iterations1 = newton.Newton(polynomial, derivative1, derivative2)
	fmt.Println("Метод Ньютона: x=", x1, " y=", y1, " Итераций=", iterations1)
}
