package main

import (
	"fmt"
	"strings"
)

func drawTable(width int, temp_list [][]float64, draw_margin func(width int), draw_line func(width int, col_1_content string, col_2_content string)) {
	draw_margin(width)

	draw_line(width, "°C", "°F")
	draw_margin(width)

	for _, temp := range temp_list {
		if len(temp) == 2 {
			draw_line(width, fmt.Sprintf("%4.2f", temp[0]), fmt.Sprintf("%4.2f", temp[1]))
		}
	}
	draw_margin(width)
}

func drawMargin(width int) {
	fmt.Println(strings.Repeat("=", width))
}

func drawLine(width int, col_1_content string, col_2_content string) {
	if width >= 7 {
		space_val := (width - 7) / 2
		format := fmt.Sprintf("| %%-%dv | %%-%dv |\n", space_val, space_val)
		fmt.Printf(format, col_1_content, col_2_content)
	}

}

func createTempList(start float64, end float64) [][]float64 {
	var temp [][]float64
	for i := start; i <= end; i += 5 {
		temp = append(temp, []float64{i, CelsiusToFahrenheit(i)})
	}
	return temp
}

func CelsiusToFahrenheit(c float64) float64 {
	return float64((c * 9.0 / 5.0) + 32.0)
}

// main():
// content := createTempList(-40, 100)
// drawTable(23, content, drawMargin, drawLine)
