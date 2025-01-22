package main

import (
	"fmt"
	"math/rand"
)

func C2FAss(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

// main
// f := K2FAss(0)
// fmt.Printf("%.2f", f)
func K2FAss(k float64) float64 {
	return ((k - 273.15) * 9.0 / 5.0) + 32
}

// Method Assignment
type kelvin float64

type celsius float64
type fahrenheit float64

func (k kelvin) KelvinToFahrenheit() fahrenheit {
	return fahrenheit(((k - 273.15) * 9.0 / 5.0) + 32)
}

func (k kelvin) KelvinToCelsius() celsius {
	return celsius(k - 273.15)
}

func (c celsius) CelsiusToFahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (c celsius) CelsiusToKelvin() kelvin {
	return kelvin(c + 273.15)
}

func (f fahrenheit) FahrenheitToCelsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func (f fahrenheit) FahrenheitToKelvin() kelvin {
	return kelvin((f-32.0)*5.0/9.0) + 273.15
}

// Closure Assiginment

type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn((151) + 150))
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	// run Method Assiginment
	var k kelvin = 294.0
	c := k.KelvinToCelsius()
	fmt.Println(c)

	// run Closure Assiginment
	var offset kelvin = 5
	sensor := calibrate(fakeSensor, 5)
	fmt.Println(sensor())
	offset++
	fmt.Println(sensor())

	// run capstone_3
	content := createTempList(-40, 100)
	drawTable(23, content, drawMargin, drawLine)
}
