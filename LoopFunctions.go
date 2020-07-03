package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	prev := 0.0
	for {
		z, prev = z-(z*z-x)/(2*z), z
		if math.Abs(prev-z) < 1e-6 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(29))
	fmt.Println(math.Sqrt(29))
}
