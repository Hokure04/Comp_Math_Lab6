package main

import (
	"bufio"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input()
}

func function(x, y float64, choice int) float64 {
	if choice == 1 {
		return y + (1+x)*math.Pow(y, 2)
	} else if choice == 2 {
		return 4*x + y/3
	}
	return 0
}

func input() {
	var y0, a, b, h, e float64
	fmt.Println("Выберите функцию для дифференцирования")
	fmt.Println("1. y + (1+x)*y^2")
	fmt.Println("2. 4*x + y/3")
	for {
		choice := bufio.NewScanner(os.Stdin)
		choice.Scan()
		input := choice.Text()

		var choiceInt int
		_, err := fmt.Sscanf(input, "%d", &choiceInt)
		if err != nil {
			fmt.Println("Ошибка: вы ввели некорректное значение")
			continue
		}

		if choiceInt > 3 || choiceInt < 1 {
			fmt.Println("Введите значение от 1 до 3")
			continue
		}
		for {
			fmt.Println("Введите начальное значение y0: ")
			y0Str := bufio.NewScanner(os.Stdin)
			y0Str.Scan()
			yInput := y0Str.Text()
			y0, err = strconv.ParseFloat(yInput, 64)
			if err != nil {
				fmt.Println("Ошибка: y0 должно быть числом")
				continue
			}
			break
		}

		for {
			fmt.Println("Введите интервал [x0, xn]: ")
			intervalStr := bufio.NewScanner(os.Stdin)
			intervalStr.Scan()
			intervalInput := intervalStr.Text()
			intervalParts := strings.Split(intervalInput, " ")
			if len(intervalParts) != 2 {
				fmt.Println("Ошибка: Интервал должен состоять из двух чисел")
				continue
			}
			a, err = strconv.ParseFloat(intervalParts[0], 64)
			if err != nil {
				fmt.Println("Ошибка: Вы ввели некоректный интервал")
				continue
			}
			b, err = strconv.ParseFloat(intervalParts[1], 64)
			if err != nil {
				fmt.Println("Ошибка: Вы ввели некорректный интервал")
				continue
			}
			break
		}

		for {
			fmt.Println("Введите значение шага h:")
			stepStr := bufio.NewScanner(os.Stdin)
			stepStr.Scan()
			stepInput := stepStr.Text()
			h, err = strconv.ParseFloat(stepInput, 64)
			if err != nil {
				fmt.Println("Ошибка: Шаг h должен быть числом")
				continue
			}
			break
		}

		for {
			fmt.Println("Введите точность: ")
			precisionStr := bufio.NewScanner(os.Stdin)
			precisionStr.Scan()
			precisionInput := precisionStr.Text()
			e, err = strconv.ParseFloat(precisionInput, 64)
			if err != nil {
				fmt.Println("Ошибка: Точность должна являться числом")
				continue
			}
			break
		}
		y1 := euler_method(y0, a, b, h, e, choiceInt)
		fmt.Printf("Решение диффиренциального уравнения по методу Эйлера: %f\n", y1)
		y2 := modified_euler(y0, a, b, h, e, choiceInt)
		fmt.Printf("Решение дифференциального уравнения по Модифицированному Эйлеру: %f\n", y2)
		milne_method(y0, a, b, h, e, choiceInt)
	}
}

func euler_method(y0, x0, xn, h, e float64, funcNumber int) float64 {
	var f float64
	var i float64
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"xi", "yi", "f(xi, yi)"})
	for i = x0; i < xn; i += h {
		f = function(i, y0, funcNumber)
		table.Append([]string{fmt.Sprintf("%f", i), fmt.Sprintf("%f", y0), fmt.Sprintf("%f", f)})
		y0 = y0 + h*f
	}
	table.Append([]string{fmt.Sprintf("%f", i), fmt.Sprintf("%f", y0)})
	table.Render()
	return y0
}

func modified_euler(y0, x0, xn, h, e float64, funcNumber int) float64 {
	var f float64
	var i float64
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"xi", "yi", "y'"})
	for i = x0; i < xn; i += h {
		f = y0 + h*function(i, y0, funcNumber)
		table.Append([]string{fmt.Sprintf("%f", i), fmt.Sprintf("%f", y0), fmt.Sprintf("%f", f)})
		y0 = y0 + h/2*(function(i, y0, funcNumber)+function(i+h, f, funcNumber))
	}
	table.Append([]string{fmt.Sprintf("%f", i), fmt.Sprintf("%f", y0)})
	table.Render()
	return y0
}

func milne_method(y0, x0, xn, h, e float64, funcNumber int) {
	var xValues []float64
	var i float64
	var condition bool
	n := (xn - x0) / h
	for i = 0; i < n; i++ {
		xValues = append(xValues, x0+h*i)
	}
	fmt.Println(xValues)
	yValues := make([]float64, int(n))
	yValues = append(yValues, y0)
	for i := 1; i < 4; i++ {
		k1 := h * function(xValues[i-1], yValues[i-1], funcNumber)
		k2 := h * function(xValues[i-1]+h/2, yValues[i-1]+k1/2, funcNumber)
		k3 := h * function(xValues[i-1]+h/2, yValues[i-1]+k2/2, funcNumber)
		k4 := h * function(xValues[i-1]+h, yValues[i-1]+k3, funcNumber)
		yValues = append(yValues, yValues[i-1]+(k1+2*k2+2*k3+k4)/6)
	}
	condition = true
	for i := 4; i < int(n); i++ {
		y := yValues[i-4] + 4*h*(2*function(xValues[i-3], yValues[i-3], funcNumber)-function(xValues[i-2], yValues[i-2], funcNumber)+2*function(xValues[i-1], yValues[i-1], funcNumber))/3
		nextY := y

		for condition {
			yc := yValues[i-2] + h*(function(xValues[i-2], yValues[i-2], funcNumber)+4*function(xValues[i-1], yValues[i-1], funcNumber)+function(xValues[i], nextY, funcNumber))/3
			if math.Abs(yc-nextY) < e {
				nextY = yc
				break
			}
			nextY = yc
		}
		yValues = append(yValues, nextY)
	}
	fmt.Println(yValues)
}
