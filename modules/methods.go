package modules

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"math"
	"os"
)

func function(x, y float64, choice int) float64 {
	if choice == 1 {
		return 4*x + y/3
	} else if choice == 2 {
		return y + math.Cos(x)
	} else if choice == 3 {
		return 1 + y + 1.5*math.Pow(x, 2)
	} else if choice == 4 {
		return y + (1+x)*math.Pow(y, 2)
	}
	return 0
}

func Euler_method(y0, x0, xn, h, e float64, funcNumber int) (float64, []float64) {
	fmt.Println("Метод эйлера:")
	var f float64
	var i float64
	var yValues []float64
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"xi", "yi", "f(xi, yi)"})
	yValues = append(yValues, y0)
	for i = x0; i < xn; i += h {
		f = function(i, y0, funcNumber)
		table.Append([]string{fmt.Sprintf("%f", i), fmt.Sprintf("%f", y0), fmt.Sprintf("%f", f)})
		y0 = y0 + h*f
		yValues = append(yValues, y0)
	}
	table.Append([]string{fmt.Sprintf("%f", i), fmt.Sprintf("%f", y0)})
	table.Render()
	fmt.Println()
	return y0, yValues
}

func Modified_euler(y0, x0, xn, h, e float64, funcNumber int) (float64, []float64) {
	fmt.Println("Модифицированный метод Эйлера:")
	var f float64
	var i float64
	var yValues []float64
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"xi", "yi", "y'"})
	yValues = append(yValues, y0)
	for i = x0; i < xn; i += h {
		f = y0 + h*function(i, y0, funcNumber)
		table.Append([]string{fmt.Sprintf("%f", i), fmt.Sprintf("%f", y0), fmt.Sprintf("%f", f)})
		y0 = y0 + h/2*(function(i, y0, funcNumber)+function(i+h, f, funcNumber))
		yValues = append(yValues, y0)
	}
	table.Append([]string{fmt.Sprintf("%f", i), fmt.Sprintf("%f", y0)})
	table.Render()
	fmt.Println()
	return y0, yValues
}

func runge_kutta_method(y0, x0, xn, h float64, funcNumber int) []float64 {
	fmt.Println("Первые 4 значения по методу Рунге 4 порядка:")
	var xValues []float64
	var j float64
	var yValues []float64
	n := math.Abs(xn-x0) / h
	for j = 0; j < n; j++ {
		xValues = append(xValues, x0+h*j)
	}
	yValues = make([]float64, 4)
	yValues[0] = y0
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"xi", "yi", "k1", "k2", "k3", "k4"})
	for i := 0; i < 3; i++ {
		k1 := h * function(xValues[i], yValues[i], funcNumber)
		k2 := h * function(xValues[i]+h/2, yValues[i]+k1/2, funcNumber)
		k3 := h * function(xValues[i]+h/2, yValues[i]+k2/2, funcNumber)
		k4 := h * function(xValues[i]+h, yValues[i]+k3, funcNumber)
		table.Append([]string{fmt.Sprintf("%f", xValues[i]), fmt.Sprintf("%f", yValues[i]), fmt.Sprintf("%f", k1), fmt.Sprintf("%f", k2), fmt.Sprintf("%f", k3), fmt.Sprintf("%f", k4)})
		yValues[i+1] = yValues[i] + (k1+2*k2+2*k3+k4)/6
	}
	table.Append([]string{fmt.Sprintf("%f", xValues[3]), fmt.Sprintf("%f", yValues[3])})
	table.Render()
	return yValues
}

func Milne_method(y0, x0, xn, h, e float64, funcNumber int) (float64, []float64) {
	fmt.Println("Метод Милна:")
	var xValues []float64
	var i float64
	n := (xn - x0) / h
	if n < 4 {
		fmt.Println("Слишком мало точек для использования метода Милна")
		return math.Inf(0), nil
	}
	for i = 0; i < n+1; i++ {
		xValues = append(xValues, x0+h*i)
	}
	yValues := make([]float64, int(n))
	yValues = runge_kutta_method(y0, x0, xn, h, funcNumber)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"xi", "predict_y", "yi"})
	for i := 4; i < len(xValues); i++ {
		predict_y := yValues[i-4] + 4*h*(2*function(xValues[i-3], yValues[i-3], funcNumber)-function(xValues[i-2], yValues[i-2], funcNumber)+2*function(xValues[i-1], yValues[i-1], funcNumber))/3
		corr_y := yValues[i-2] + h*(function(xValues[i-2], yValues[i-2], funcNumber)+4*function(xValues[i-1], yValues[i-1], funcNumber)+function(xValues[i], predict_y, funcNumber))/3
		for math.Abs(predict_y-corr_y) > e {
			predict_y = corr_y
			corr_y = yValues[i-2] + h*(function(xValues[i-2], yValues[i-2], funcNumber)+4*function(xValues[i-1], yValues[i-1], funcNumber)+function(xValues[i], predict_y, funcNumber))/3
		}
		yValues = append(yValues, corr_y)
		table.Append([]string{fmt.Sprintf("%f", xValues[i]), fmt.Sprintf("%f", corr_y), fmt.Sprintf("%f", yValues[i])})
	}
	table.Render()
	y := yValues[int(n)]
	fmt.Println()
	return y, yValues
}
