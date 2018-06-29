package main

import (
	"fmt"
	"benben"
)

func main() {
	math.Add(1, 2)
	/*
	x := 1
	y := 2
	fmt.Printf("x: %d, y: %d", x, y)
	x, y = y, x
	fmt.Printf("\nx: %d, y: %d", x, y)

	type Celsius float64
	type Fahrenheit float64

	const (
		AbsoluteZeroC Celsius = -2
	)
	*/
	/*
	fmt.Printf("x: %d, y: %d", x, y)
	const float, boilingF = 32.0, 212.0
	fmt.Printf("%g F = %g C\n", float, fToC(float))
	fmt.Printf("%g F = %g C\n", boilingF, fToC(boilingF))
	*/
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}