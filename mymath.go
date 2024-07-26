package mymath

import "math"

func Abs(x float64) float64 {
	return math.Abs(x)
}

func Yn(n int, x float64) float64 {
	return math.Yn(n, x)
}

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

func Ceil(x float64) float64 {
	return math.Ceil(x)
}

func Floor(x float64) float64 {
	return math.Floor(x)
}

func Min(x, y float64) float64 {
	return math.Min(x, y)
}

func Max(x, y float64) float64 {
	return math.Max(x, y)
}

func Trunc(x float64) float64 {
	return math.Trunc(x)
}