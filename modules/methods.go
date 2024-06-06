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

func true_value_calc(x float64, choice int) float64 {
	if choice == 1 {
		return math.Exp(x/3) - 12*x - 36
	} else if choice == 2 {
		return math.Sin(x)/2 - math.Cos(x)/2 + math.Exp(x)
	} else if choice == 3 {
		return math.Exp(x) - 3*math.Pow(x, 2)/2 - 3*x - 4
	} else if choice == 4 {
		return -math.Exp(x) / (x*math.Exp(x) + 1)
	}
	return 0
}

func Euler_method(y0, h, e float64, funcNumber int, xValues []float64) []float64 {
	fmt.Println("Метод эйлера:")
	var f float64
	var yValues []float64
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"i", "xi", "yi", "f(xi, yi)"})

	for i, x := range xValues {
		yValues = append(yValues, y0)
		f = function(x, y0, funcNumber)
		table.Append([]string{fmt.Sprintf("%d", i), fmt.Sprintf("%f", x), fmt.Sprintf("%f", y0), fmt.Sprintf("%f", f)})
		y0 = y0 + h*f
	}
	table.Render()
	fmt.Println()
	return yValues
}

func Modified_euler(y0, h, e float64, funcNumber int, xValues []float64) []float64 {
	fmt.Println("Модифицированный метод Эйлера:")
	var f float64
	var yValues []float64
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"i", "xi", "yi", "y'"})
	for i, x := range xValues {
		yValues = append(yValues, y0)
		f = y0 + h*function(x, y0, funcNumber)
		table.Append([]string{fmt.Sprintf("%d", i), fmt.Sprintf("%f", x), fmt.Sprintf("%f", y0), fmt.Sprintf("%f", f)})
		y0 = y0 + h/2*(function(x, y0, funcNumber)+function(x+h, f, funcNumber))

	}
	table.Render()
	fmt.Println()
	return yValues
}

func runge_kutta_method(y0, h float64, funcNumber int, xValues []float64) []float64 {
	fmt.Println("Первые 4 значения по методу Рунге 4 порядка:")
	var yValues []float64

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

func Milne_method(y0, h, e float64, funcNumber int, xValues []float64) ([]float64, []float64) {
	fmt.Println("Метод Милна:")
	var true_y_values []float64
	if len(xValues) < 4 {
		fmt.Println("Слишком мало точек для использования метода Милна")
		return nil, nil
	}
	for i := 0; i < len(xValues); i++ {
		true_y_values = append(true_y_values, true_value_calc(xValues[i], funcNumber))
	}
	yValues := make([]float64, len(xValues))
	yValues = runge_kutta_method(y0, h, funcNumber, xValues)
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
	fmt.Println()
	return yValues, true_y_values
}
