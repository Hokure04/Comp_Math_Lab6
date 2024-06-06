package main

import (
	"Comp_Math_Lab6/modules"
	"bufio"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	input()
}

func input() {
	var y0, a, b, h, e float64
	for {
		fmt.Println("Выберите функцию для дифференцирования")
		fmt.Println("1. 4*x + y/3")
		fmt.Println("2. y + cos(x)")
		fmt.Println("3. 1 + y + 1.5*x^2")
		fmt.Println("4. y + (1 + x) * y^2")
		choice := bufio.NewScanner(os.Stdin)
		choice.Scan()
		input := choice.Text()

		var choiceInt int
		_, err := fmt.Sscanf(input, "%d", &choiceInt)
		if err != nil {
			fmt.Println("Ошибка: вы ввели некорректное значение")
			continue
		}

		if choiceInt > 4 || choiceInt < 1 {
			fmt.Println("Введите значение от 1 до 4")
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
			if a >= b {
				fmt.Println("Ошибка: Первое значение должно быть меньше второго")
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
			if h <= 0 {
				fmt.Println("Шаг должен являться положительным числом отличным от нуля")
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
			if e <= 0 {
			}
			fmt.Println("Ошибка: Точность должна являться положительным числом отличным от нуля")
			break
		}
		y1, yValues1 := modules.Euler_method(y0, a, b, h, e, choiceInt)
		fmt.Printf("Решение диффиренциального уравнения по методу Эйлера: %f\n", y1)
		y2, yValues2 := modules.Modified_euler(y0, a, b, h, e, choiceInt)
		fmt.Printf("Решение дифференциального уравнения по Модифицированному Эйлеру: %f\n", y2)
		y3, yValues3 := modules.Milne_method(y0, a, b, h, e, choiceInt)

		if y3 == math.Inf(0) && yValues3 == nil {
			continue
		}
		fmt.Printf("Решение дифференциального уравнения по методу Милна: %f\n", y3)

		xValues := make([]float64, len(yValues1))
		for i := 0; i < len(yValues1); i++ {
			xValues[i] = a + h*float64(i)
		}

		plotGraphs(yValues1, yValues2, yValues3, xValues)
	}
}

func plotGraphs(yValues1, yValues2, yValues3 []float64, xValues []float64) {
	dirPath := "graphs"
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := plot.New()

	p.Title.Text = "Графики методов Эйлера, Модифицированного Эйлера и Милна"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	eulerPoints := make(plotter.XYs, len(yValues1))
	for i, y := range yValues1 {
		eulerPoints[i].X = xValues[i]
		eulerPoints[i].Y = y
	}
	eulerLine, err := plotter.NewLine(eulerPoints)
	if err != nil {
		panic(err)
	}
	eulerLine.Color = color.RGBA{R: 255, A: 255}
	p.Add(eulerLine)

	modifiedEulerPoints := make(plotter.XYs, len(yValues2))
	for i, y := range yValues2 {
		modifiedEulerPoints[i].X = xValues[i]
		modifiedEulerPoints[i].Y = y
	}
	modifiedEulerLine, err := plotter.NewLine(modifiedEulerPoints)
	if err != nil {
		panic(err)
	}
	modifiedEulerLine.Color = color.RGBA{G: 255, A: 255}
	p.Add(modifiedEulerLine)

	milnePoints := make(plotter.XYs, len(yValues3))
	for i, y := range yValues3 {
		milnePoints[i].X = xValues[i]
		milnePoints[i].Y = y
	}
	milneLine, err := plotter.NewLine(milnePoints)
	if err != nil {
		panic(err)
	}
	milneLine.Color = color.RGBA{B: 255, A: 255}
	p.Add(milneLine)

	fileIndex := 1
	fileName := fmt.Sprintf("methods_graphs_%d.png", fileIndex)
	filePath := filepath.Join(dirPath, fileName)
	for {
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			break
		}
		fileIndex++
		fileName = fmt.Sprintf("methods_graphs_%d.png", fileIndex)
		filePath = filepath.Join(dirPath, fileName)
	}

	err = p.Save(6*vg.Inch, 4*vg.Inch, filePath)
	if err != nil {
		panic(err)
	}
}
