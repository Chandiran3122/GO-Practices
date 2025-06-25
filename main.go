package main

import "fmt"

func avg(x float64, y float64) float64 {
	return (x + y) / 2
}

func main() {
	p := 23.34
	q := 36.25

	result := avg(p, q)

	fmt.Printf("Average of %.2f and %.2f = %.2f\n", p, q, result)
}
